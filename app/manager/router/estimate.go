package router

import (
	"sync"

	"github.com/openPanel/core/app/clients/http"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

func EstimateAndStoreLatencies() map[Edge]int {
	nodesLock.RLock()

	wg := sync.WaitGroup{}
	infos := map[Edge]int{}

	for _, node := range nodes {
		if node.Id == global.App.NodeInfo.ServerId {
			continue
		}

		wg.Add(1)
		go func(node Node) {
			defer wg.Done()

			latency, err := http.TcpPing(node.AddrPort)
			if err != nil {
				log.Debugf("failed to ping node %s: %v", node.AddrPort.String(), err)
				return
			}

			infos[Edge{
				From: global.App.NodeInfo.ServerId,
				To:   node.Id,
			}] = latency
		}(node)
	}

	wg.Wait()
	log.Debugf("estimated latencies: %v", infos)

	nodesLock.RUnlock()

	UpdateRouterInfo(infos)

	return infos
}

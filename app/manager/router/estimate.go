package router

import (
	"sync"

	"github.com/openPanel/core/app/clients/http"
	"github.com/openPanel/core/app/global"
)

func EstimateLatencies() map[Edge]int {
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

			latency, err := http.QuicPing(node.AddrPort)
			if err != nil {
				return
			}

			infos[Edge{
				From: global.App.NodeInfo.ServerId,
				To:   node.Id,
			}] = latency
		}(node)
	}

	wg.Wait()

	nodesLock.RUnlock()

	UpdateRouterInfo(infos)

	return infos
}

package router

import (
	"fmt"
	"sync"

	"github.com/openPanel/core/app/clients/http"
	"github.com/openPanel/core/app/global"
)

var EstimateLatenciesCallback func(map[Edge]int) = nil

func estimateLatencies() {
	nodesLock.Lock()
	defer nodesLock.Unlock()

	wg := sync.WaitGroup{}
	infos := map[Edge]int{}

	for _, node := range nodes {
		if node.Id == global.App.NodeInfo.ServerId {
			continue
		}

		wg.Add(1)
		go func(node Node) {
			defer wg.Done()

			latency, err := http.QuicPing(fmt.Sprintf("https://%s:%d", node.Address.Ip, node.Address.Port))
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
	updateRouterInfo(infos)
	updateRouterDecision()

	if EstimateLatenciesCallback != nil {
		EstimateLatenciesCallback(infos)
	}
}

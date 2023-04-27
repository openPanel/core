package router

import (
	"net/netip"
	"sync"

	"github.com/openPanel/core/app/clients/tcp"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

func EstimateLatencies(nodes []Node) map[Edge]int {
	wg := sync.WaitGroup{}
	infos := map[Edge]int{}

	for _, node := range nodes {
		wg.Add(1)
		go func(id string, addr netip.AddrPort) {
			defer wg.Done()

			latency, err := tcp.Ping(addr)
			if err != nil {
				log.Debugf("failed to ping node %s: %v", addr.String(), err)
				return
			}

			infos[Edge{
				From: global.App.NodeInfo.ServerId,
				To:   id,
			}] = latency
		}(node.Id, node.AddrPort)
	}

	wg.Wait()
	log.Debugf("estimated latencies: %v", infos)

	return infos
}

// EstimateAndStoreLatencies Estimate the latency between nodes and store it in the router info
func EstimateAndStoreLatencies() map[Edge]int {
	newNodes := make([]Node, 0, len(nodes) - 1)
	ndLock.RLock()
	for _, node := range flattenNodes() {
		if node.Id != global.App.NodeInfo.ServerId {
			newNodes = append(newNodes, node)
		}
	}
	ndLock.RUnlock()

	infos := EstimateLatencies(newNodes)

	UpdateLinkStates(infos)

	return infos
}

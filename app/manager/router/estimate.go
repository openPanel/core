package router

import (
	"net/netip"
	"sync"

	"github.com/openPanel/core/app/clients/tcp"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

func EstimateLatencies(nodes []Node, from string) LinkStates {
	wg := sync.WaitGroup{}
	wg.Add(len(nodes))
	infos := map[Edge]int{}

	for _, node := range nodes {
		go func(id string, addr netip.AddrPort) {
			defer wg.Done()

			latency, err := tcp.Ping(addr)
			if err != nil {
				log.Debugf("failed to ping node %s: %v", addr.String(), err)
				return
			}

			infos[Edge{
				From: from,
				To:   id,
			}] = latency
		}(node.Id, node.AddrPort)
	}

	wg.Wait()
	log.Debugf("estimated latencies: %v", infos)

	return infos
}

// EstimateAndStoreLatencies Estimate the latency between nodes and store it in the router info
func EstimateAndStoreLatencies() LinkStates {
	filteredNodes := make([]Node, 0, len(nodes)-1)
	ndLock.RLock()
	for _, node := range FlattedNodes() {
		if node.Id != global.App.NodeInfo.ServerId {
			filteredNodes = append(filteredNodes, node)
		}
	}
	ndLock.RUnlock()

	infos := EstimateLatencies(filteredNodes, global.App.NodeInfo.ServerId)

	UpdateLinkStates(infos)

	return infos
}

package router

import (
	"math"
	"net/netip"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

// implement a default route algorithm
// It directly connects to other node if it is possible
// For unreachable node, it will use the lowest latency node to connect to it
// This algorithm will be used before received all node link state.
func defaultRouteAlgorithm() {
	if linkStates == nil || len(linkStates) == 0 {
		log.Warnf("linkStates is empty, can not generate router decisions")
		return
	}

	// reset router decision
	routerDecisions = map[string]netip.AddrPort{}

	// find the lowest latency node
	var lowestLatencyNode string
	var lowestLatency = math.MaxInt32

	var reachableNodes = map[string]int{}

	// find all reachable node
	for edge, latency := range linkStates {
		if edge.From != global.App.NodeInfo.ServerId {
			continue
		}

		if latency < lowestLatency {
			lowestLatency = latency
			lowestLatencyNode = edge.To
		}

		reachableNodes[edge.To] = latency
	}

	for id, node := range nodes {
		if id == global.App.NodeInfo.ServerId {
			continue
		}
		if _, ok := reachableNodes[id]; ok {
			routerDecisions[id] = node
		} else {
			routerDecisions[id] = nodes[lowestLatencyNode]
		}
	}
}

package router

import (
	"math"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

// implement a default route algorithm
// It directly connects to other node if it is possible
// For unreachable node, it will use the lowest latency node to connect to it
// This algorithm will be used before received all node link state.
func defaultRouteAlgorithm() {
	if routerInfos == nil || len(routerInfos) == 0 {
		log.Warnf("routerInfos is empty, can not generate router decisions")
		return
	}

	// reset router decision
	routerDecision = map[string]Address{}

	// find the lowest latency node
	var lowestLatencyNode string
	var lowestLatency = math.MaxInt32

	var reachableNodes = map[string]int{}

	// find all reachable node
	for edge, latency := range routerInfos {
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
			routerDecision[id] = node.Address
		} else {
			routerDecision[id] = nodes[lowestLatencyNode].Address
		}
	}
}

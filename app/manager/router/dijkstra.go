package router

import (
	"container/heap"
	"math"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

func dijkstraRouteAlgorithm() {
	distances := make(map[string]int)
	previous := make(map[string]string)

	for id := range nodes {
		distances[id] = math.MaxInt32
		previous[id] = ""
	}

	startNodeId := global.App.NodeInfo.ServerId
	distances[startNodeId] = 0

	Q := make(priorityQueue, 0)
	heap.Push(&Q, &pqNode{id: startNodeId, distance: 0})

	for Q.Len() > 0 {
		u := heap.Pop(&Q).(*pqNode)
		for edge, latency := range routerInfos {
			if edge.From != u.id {
				continue
			}
			alt := distances[u.id] + latency
			if alt < distances[edge.To] {
				distances[edge.To] = alt
				previous[edge.To] = u.id
				heap.Push(&Q, &pqNode{id: edge.To, distance: alt})
			}
		}
	}

	var calPath func(nodeId string) []string
	calPath = func(nodeId string) []string {
		if len(previous[nodeId]) == 0 {
			return []string{nodeId}
		}
		var path = make([]string, 0)
		path = append(path, calPath(previous[nodeId])...)
		path = append(path, nodeId)
		return path
	}

	for _, node := range nodes {
		if node.Id == global.App.NodeInfo.ServerId {
			continue
		}
		path := calPath(node.Id)
		if len(path) < 2 {
			log.Warnf("No path to node %s", node.Id)
		}
		routerDecision[node.Id] = nodes[path[1]].Address
	}
}

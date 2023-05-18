package router

import (
	"math"

	pq "github.com/ugurcsen/gods-generic/queues/priorityqueue"
	"github.com/ugurcsen/gods-generic/utils"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

type netNode struct {
	id       string
	distance int
}

func byDistance(a, b netNode) int {
	return utils.NumberComparator(a.distance, b.distance)
}

func dijkstraRouteAlgorithm() {
	distances := make(map[string]int, len(nodes))
	previous := make(map[string]string, len(nodes))

	for id := range nodes {
		distances[id] = math.MaxInt32
		previous[id] = ""
	}

	startNodeId := global.App.NodeInfo.ServerId
	distances[startNodeId] = 0

	Q := pq.NewWith(byDistance)
	Q.Enqueue(netNode{id: startNodeId, distance: 0})

	for !Q.Empty() {
		u, _ := Q.Dequeue()

		if distances[u.id] < u.distance {
			continue
		}

		for _, edge := range linkStates[u.id] {
			alt := distances[u.id] + edge.Latency
			if alt < distances[edge.To] {
				distances[edge.To] = alt
				previous[edge.To] = u.id
				Q.Enqueue(netNode{id: edge.To, distance: alt})
			}
		}
	}

	var calPath func(nodeId string) []string
	calPath = func(nodeId string) []string {
		if len(previous[nodeId]) == 0 {
			return []string{nodeId}
		}
		var path = make([]string, 0, 6)
		path = append(path, calPath(previous[nodeId])...)
		path = append(path, nodeId)
		return path
	}

	for id := range nodes {
		if id == global.App.NodeInfo.ServerId {
			continue
		}
		path := calPath(id)
		if len(path) < 2 {
			log.Warnf("No path to node %s", id)
		}
		decisions[id] = nodes[path[1]]
	}
}

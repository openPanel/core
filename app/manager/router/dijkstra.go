//go:build !pq

package router

import (
	"math"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

type netNode struct {
	id       string
	distance int
}

func dijkstraRouteAlgorithm() {
	distances := make(map[string]int, len(nodes))
	previous := make(map[string]string, len(nodes))
	visited := make(map[string]bool, len(nodes))

	for id := range nodes {
		distances[id] = math.MaxInt32
		previous[id] = ""
	}

	startNodeId := global.App.NodeInfo.ServerId
	distances[startNodeId] = 0

	for len(visited) < len(nodes) {
		// Find the unvisited node with the smallest distance
		minDistance := math.MaxInt32
		minNode := netNode{}

		for id, distance := range distances {
			if !visited[id] && distance < minDistance {
				minDistance = distance
				minNode = netNode{id: id, distance: distance}
			}
		}

		visited[minNode.id] = true

		for _, edge := range linkStates[minNode.id] {
			alt := distances[minNode.id] + edge.Latency
			if alt < distances[edge.To] {
				distances[edge.To] = alt
				previous[edge.To] = minNode.id
			}
		}
	}

	var resolve func(nodeId string) []string
	resolve = func(nodeId string) []string {
		if len(previous[nodeId]) == 0 {
			return []string{nodeId}
		}
		var path = make([]string, 0, 6)
		path = append(path, resolve(previous[nodeId])...)
		path = append(path, nodeId)
		return path
	}

	for id := range nodes {
		if id == global.App.NodeInfo.ServerId {
			continue
		}
		path := resolve(id)
		if len(path) < 2 {
			log.Warnf("No path to node %s", id)
		}
		decisions[id] = nodes[path[1]]
	}
}

package router

import (
	"fmt"
	"net/netip"
	"sync"

	"github.com/pkg/errors"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

type Edge struct {
	From string `json:"f"`
	To   string `json:"t"`
}

type LinkStates = map[Edge]int

// RouterInfo The map store the latency between two nodes
var linkStates = LinkStates{}
var lsLock = sync.RWMutex{}

// nodes The map store the nodes info
var nodes = map[string]netip.AddrPort{}
var ndLock = sync.RWMutex{}

// RouterDecisions The map store the decision of the router, value define the next hop
var routerDecisions = make(map[string]netip.AddrPort)
var rdLock = sync.RWMutex{}

type Node struct {
	Id       string
	AddrPort netip.AddrPort
}

// NodePersistence just prevent import cycle
var NodePersistence func([]Node)

func FlattedNodes() []Node {
	flat := make([]Node, 0, len(nodes))
	for id, addr := range nodes {
		flat = append(flat, Node{
			Id:       id,
			AddrPort: addr,
		})
	}
	return flat
}

func AddNodes(ns []Node) {
	ndLock.Lock()
	defer ndLock.Unlock()
	for _, node := range ns {
		nodes[node.Id] = node.AddrPort
	}

	if NodePersistence != nil {
		NodePersistence(FlattedNodes())
	}
}

func SetNodes(ns []Node) {
	ndLock.Lock()
	defer ndLock.Unlock()
	nodes = make(map[string]netip.AddrPort, len(ns))
	for _, node := range ns {
		nodes[node.Id] = node.AddrPort
	}

	if NodePersistence != nil {
		NodePersistence(FlattedNodes())
	}

	filterLinkStates("")
}

func EditNodes(op func(ns []Node) []Node) {
	ndLock.Lock()
	defer ndLock.Unlock()

	SetNodes(op(FlattedNodes()))
}

// Invalid all outdated link states
// Delete all router info that related to the node which id is given
func filterLinkStates(id string) {
	lsLock.Lock()
	defer lsLock.Unlock()

	// we assume nodes has already been locked
	nodeMap := make(map[string]bool)
	for nodeId := range nodes {
		nodeMap[nodeId] = true
	}

	if len(linkStates) == 0 {
		return
	}

	for link := range linkStates {
		if link.From == id || link.To == id {
			delete(linkStates, link)
		}
		// remove no exist router info
		if !nodeMap[link.From] || !nodeMap[link.To] {
			delete(linkStates, link)
		}
	}
}

func UpdateLinkStates(infos ...LinkStates) {
	lsLock.Lock()

	for _, info := range infos {
		for link, latency := range info {
			linkStates[link] = latency
		}
	}

	lsLock.Unlock()

	calculateRouting()
}

func GetLinkStates() LinkStates {
	lsLock.RLock()
	defer lsLock.RUnlock()
	return linkStates
}

func canUseDijkstra() bool {
	ndLock.RLock()
	defer ndLock.RUnlock()

	currentNode := global.App.NodeInfo.ServerId

	// dfs to find all nodes that can be reached from current node
	visited := make(map[string]bool)
	var dfs func(string)
	dfs = func(node string) {
		visited[node] = true
		for edge := range linkStates {
			if edge.From == node && !visited[edge.To] {
				dfs(edge.To)
			}
		}
	}

	dfs(currentNode)

	return len(visited) == len(nodes)
}

func calculateRouting() {
	rdLock.Lock()
	defer rdLock.Unlock()
	lsLock.RLock()
	defer lsLock.RUnlock()
	ndLock.RLock()
	defer ndLock.RUnlock()

	log.Debugf("link states: %v", linkStates)
	log.Debugf("nodes: %v", nodes)

	if canUseDijkstra() {
		log.Debugf("all nodes have been connected, use dijkstra route algorithm")
		// dijkstra needs info of all nodes
		dijkstraRouteAlgorithm()
	} else {
		log.Debugf("not all nodes have been connected, use default route algorithm")
		defaultRouteAlgorithm()
	}

	log.Debugf("router decisions: %v", routerDecisions)
}

func GetHop(id string) (netip.AddrPort, error) {
	rdLock.RLock()
	defer rdLock.RUnlock()
	addr, ok := routerDecisions[id]
	if !ok {
		log.Debugf("no route to %s, %v", id, routerDecisions)
		return netip.AddrPort{}, errors.New(fmt.Sprintf("no route to %s", id))
	}
	return addr, nil
}

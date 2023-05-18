package router

import (
	"fmt"
	"net/netip"
	"sync"

	"github.com/pkg/errors"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
)

type LinkState struct {
	From    string `json:"u"`
	To      string `json:"v"`
	Latency int    `json:"w"`
}

type LinkStates = []LinkState

// RouterInfo The map store the latency between two nodes
var linkStates = map[string][]LinkState{}
var lsLock = sync.RWMutex{}

// nodes The map store the nodes info
var nodes = map[string]netip.AddrPort{}
var ndLock = sync.RWMutex{}

// RouterDecisions The map store the decision of the router, value define the next hop
var decisions = make(map[string]netip.AddrPort)
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
}

func EditNodes(op func(ns []Node) []Node) {
	ndLock.Lock()
	defer ndLock.Unlock()

	SetNodes(op(FlattedNodes()))
}

// UpdateLinkStates Update the latency between nodes
// All link from the node included in the infos will be cleared before update
func UpdateLinkStates(infos ...LinkStates) {
	lsLock.Lock()

	tmp := make(map[string]map[string]int)

	editedNodes := make(map[string]bool)

	for _, info := range infos {
		for _, link := range info {
			if _, ok := tmp[link.From]; !ok {
				tmp[link.From] = make(map[string]int)
			}
			tmp[link.From][link.To] = link.Latency
			editedNodes[link.From] = true
		}
	}

	for node := range editedNodes {
		linkStates[node] = make([]LinkState, 0, len(tmp[node]))
	}

	for node, links := range tmp {
		for to, latency := range links {
			linkStates[node] = append(linkStates[node], LinkState{
				From:    node,
				To:      to,
				Latency: latency,
			})
		}
	}

	lsLock.Unlock()

	calculateRouting()
}

func GetLinkStates() LinkStates {
	lsLock.RLock()
	defer lsLock.RUnlock()
	ret := make(LinkStates, 0)
	for _, links := range linkStates {
		ret = append(ret, links...)
	}
	return ret
}

func canUseDijkstra() bool {
	ndLock.RLock()
	defer ndLock.RUnlock()

	currentNode := global.App.NodeInfo.ServerId

	// dfs to check if all nodes that can be reached from current node
	visited := make(map[string]bool)
	var dfs func(string)
	dfs = func(node string) {
		visited[node] = true
		for _, link := range linkStates[node] {
			if !visited[link.To] {
				dfs(link.To)
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

	log.Debugf("router decisions: %v", decisions)
}

func GetHop(id string) (netip.AddrPort, error) {
	rdLock.RLock()
	defer rdLock.RUnlock()
	addr, ok := decisions[id]
	if !ok {
		log.Debugf("no route to %s, %v", id, decisions)
		return netip.AddrPort{}, errors.New(fmt.Sprintf("no route to %s", id))
	}
	return addr, nil
}

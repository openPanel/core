package router

import (
	"net"
	"net/netip"
	"sync"

	"github.com/pkg/errors"
)

type Edge struct {
	From string
	To   string
}

// RouterInfo The map store the latency between two nodes
var routerInfos = map[Edge]int{}
var riLock = sync.RWMutex{}

var nodes = map[string]Node{}
var nodesLock = sync.RWMutex{}

// RouterDecision The map store the decision of the router, value define the next hop
var routerDecision = make(map[string]netip.AddrPort)
var rdLock = sync.RWMutex{}

type Node struct {
	Id       string
	AddrPort netip.AddrPort
}

// NodePersistence just prevent import cycle
var NodePersistence func([]Node)

func flattenNodes() []Node {
	flat := make([]Node, 0, len(nodes))
	for _, node := range nodes {
		flat = append(flat, node)
	}
	return flat
}

func AddNodes(ns []Node) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	for _, node := range ns {
		nodes[node.Id] = node
	}

	if NodePersistence != nil {
		NodePersistence(flattenNodes())
	}
}

func SetNodes(ns []Node) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	nodes = make(map[string]Node)
	for _, node := range ns {
		nodes[node.Id] = node
	}

	if NodePersistence != nil {
		NodePersistence(flattenNodes())
	}

	filterRouterInfos("")
}

func DeleteNode(id string) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	delete(nodes, id)

	if NodePersistence != nil {
		NodePersistence(flattenNodes())
	}

	filterRouterInfos(id)
}

func UpdateNode(id string, ip net.IP, port int) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	nodes[id] = Node{
		Id:       id,
		AddrPort: netip.AddrPortFrom(netip.MustParseAddr(ip.String()), uint16(port)),
	}

	filterRouterInfos(id)
}

// Invalid all outdated router info
func filterRouterInfos(id string) {
	riLock.Lock()
	defer riLock.Unlock()

	// we assume nodes has already been locked
	nodeMap := make(map[string]bool)
	for _, node := range nodes {
		nodeMap[node.Id] = true
	}

	if len(routerInfos) == 0 {
		return
	}

	for link := range routerInfos {
		if link.From == id || link.To == id {
			delete(routerInfos, link)
		}
		// remove no exist router info
		if !nodeMap[link.From] || !nodeMap[link.To] {
			delete(routerInfos, link)
		}
	}
}

func UpdateRouterInfo(infos map[Edge]int) {
	riLock.Lock()

	for link, latency := range infos {
		routerInfos[link] = latency
	}

	riLock.Unlock()

	updateRouterDecision()
}

func GetRouterInfo() map[Edge]int {
	riLock.RLock()
	defer riLock.RUnlock()
	return routerInfos
}

func updateRouterDecision() {
	rdLock.Lock()
	defer rdLock.Unlock()
	riLock.RLock()
	defer riLock.RUnlock()
	nodesLock.RLock()
	defer nodesLock.RUnlock()

	var nodeIds = make(map[string]bool)
	for edge := range routerInfos {
		nodeIds[edge.From] = true
	}

	if len(nodeIds) <= 2 {
		return
	}

	if len(nodeIds) < len(nodes) {
		defaultRouteAlgorithm()
	} else {
		// dijkstra needs info of all nodes
		dijkstraRouteAlgorithm()
	}
}

func GetHop(id string) (netip.AddrPort, error) {
	rdLock.RLock()
	defer rdLock.RUnlock()
	addr, ok := routerDecision[id]
	if !ok {
		return netip.AddrPort{}, errors.New("no route to host")
	}
	return addr, nil
}

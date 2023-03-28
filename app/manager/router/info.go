package router

import (
	"net"
	"sync"
)

type Edge struct {
	From string
	To   string
}

// RouterInfo The map store the latency between two nodes
var routerInfos = map[Edge]int{}
var riLock = sync.RWMutex{}

type Address struct {
	Ip   net.IP
	Port int
}

var nodes map[string]Node
var nodesLock = sync.RWMutex{}

// RouterDecision The map store the decision of the router, value define the next hop
var routerDecision = make(map[string]Address)
var rdLock = sync.RWMutex{}

type Node struct {
	Id string
	Address
}

func AddNodes(ns []Node) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	for _, node := range ns {
		nodes[node.Id] = node
	}
}

func DeleteNode(id string) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	delete(nodes, id)

	filterRouterInfos(id)
}

func UpdateNode(id string, ip net.IP, port int) {
	nodesLock.Lock()
	defer nodesLock.Unlock()
	nodes[id] = Node{
		Id: id,
		Address: Address{
			Ip:   ip,
			Port: port,
		},
	}

	filterRouterInfos(id)
}

// Invalid all outdated router info
func filterRouterInfos(id string) {
	riLock.Lock()
	defer riLock.Unlock()

	if len(routerInfos) == 0 {
		return
	}

	for link := range routerInfos {
		if link.From == id || link.To == id {
			delete(routerInfos, link)
		}
	}
}

func updateRouterInfo(infos map[Edge]int) {
	riLock.Lock()
	defer riLock.Unlock()

	for link, latency := range infos {
		routerInfos[link] = latency
	}
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

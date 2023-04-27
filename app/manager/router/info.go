package router

import (
	"net"
	"net/netip"
	"sync"

	"github.com/pkg/errors"

	"github.com/openPanel/core/app/tools/utils/netUtils"
)

type Edge struct {
	From string
	To   string
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

func flattenNodes() []Node {
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
		NodePersistence(flattenNodes())
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
		NodePersistence(flattenNodes())
	}

	filterRouterInfos("")
}

func DeleteNode(id string) {
	ndLock.Lock()
	defer ndLock.Unlock()
	delete(nodes, id)

	if NodePersistence != nil {
		NodePersistence(flattenNodes())
	}

	filterRouterInfos(id)
}

func UpdateNode(id string, ip net.IP, port int) {
	ndLock.Lock()
	defer ndLock.Unlock()
	nodes[id] = netUtils.NewAddrPortWithIP(ip, port)

	filterRouterInfos(id)
}

// Invalid all outdated router info
// Delete all router info that related to the node which id is given
func filterRouterInfos(id string) {
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

func UpdateLinkStates(infos LinkStates) {
	lsLock.Lock()

	for link, latency := range infos {
		linkStates[link] = latency
	}

	lsLock.Unlock()

	updateRouterDecision()
}

func GetLinkStates() LinkStates {
	lsLock.RLock()
	defer lsLock.RUnlock()
	return linkStates
}

func updateRouterDecision() {
	rdLock.Lock()
	defer rdLock.Unlock()
	lsLock.RLock()
	defer lsLock.RUnlock()
	ndLock.RLock()
	defer ndLock.RUnlock()

	var nodeIds = make(map[string]bool)
	for edge := range linkStates {
		nodeIds[edge.From] = true
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
	addr, ok := routerDecisions[id]
	if !ok {
		return netip.AddrPort{}, errors.New("no route to host")
	}
	return addr, nil
}

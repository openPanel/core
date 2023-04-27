package bootstrap

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

func createEmptyNetGraph() {
	nodes := []router.Node{
		{
			Id: global.App.NodeInfo.ServerId,
			AddrPort: netUtils.NewAddrPortWithIP(
				global.App.NodeInfo.ServerPublicIP,
				global.App.NodeInfo.ServerPort,
			),
		},
	}
	router.AddNodes(nodes)
}

// loadAndSaveInitialNodes return a list of nodes, current node at index 0
func loadAndSaveInitialNodes(nodes []*pb.Node, self router.Node) []router.Node {
	routerNodes := make([]router.Node, len(nodes)+1)
	routerNodes[0] = self
	for i, node := range nodes {
		routerNodes[i+1] = router.Node{
			Id:       node.Id,
			AddrPort: netUtils.NewAddrPortWithString(node.Ip, int(node.Port)),
		}
	}

	router.SetNodes(routerNodes)

	return routerNodes
}

func loadLinkStates(lst []*pb.LinkState) {

	edges := make(map[router.Edge]int)
	for _, state := range lst {
		edges[router.Edge{
			From: state.From,
			To:   state.To,
		}] = int(state.Latency)
	}

	router.UpdateLinkStates(edges)

	router.EstimateAndStoreLatencies()
}

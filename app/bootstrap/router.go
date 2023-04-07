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

func createFullNetGraphAtJoin(resp *pb.RegisterResponse) {
	nodes := make([]router.Node, len(resp.Nodes))
	for i, node := range resp.Nodes {
		nodes[i] = router.Node{
			Id:       node.Id,
			AddrPort: netUtils.NewAddPortWithString(node.Ip, int(node.Port)),
		}
	}

	router.AddNodes(nodes)

	edges := make(map[router.Edge]int)
	for _, edge := range resp.LinkStates {
		edges[router.Edge{
			From: edge.From,
			To:   edge.To,
		}] = int(edge.Latency)
	}

	router.UpdateRouterInfo(edges)

	router.EstimateAndStoreLatencies()
}

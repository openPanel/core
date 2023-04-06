package bootstrap

import (
	"net/netip"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/router"
)

func createEmptyNetGraph() {
	nodes := []router.Node{
		{
			Id: global.App.NodeInfo.ServerId,
			AddrPort: netip.AddrPortFrom(
				netip.MustParseAddr(global.App.NodeInfo.ServerPublicIP.String()),
				uint16(global.App.NodeInfo.ServerPort),
			),
		},
	}
	router.AddNodes(nodes)
}

func createFullNetGraphAtJoin(resp *pb.RegisterResponse) {
	nodes := make([]router.Node, len(resp.Nodes)+1)
	for i, node := range resp.Nodes {
		nodes[i] = router.Node{
			Id:       node.Id,
			AddrPort: netip.AddrPortFrom(netip.MustParseAddr(node.Ip), uint16(node.Port)),
		}
	}

	nodes = append(nodes, router.Node{
		Id: global.App.NodeInfo.ServerId,
		AddrPort: netip.AddrPortFrom(
			netip.MustParseAddr(global.App.NodeInfo.ServerPublicIP.String()),
			uint16(global.App.NodeInfo.ServerPort),
		),
	})

	router.AddNodes(nodes)

	edges := make(map[router.Edge]int)
	for _, edge := range resp.LinkStates {
		edges[router.Edge{
			From: edge.From,
			To:   edge.To,
		}] = int(edge.Latency)
	}

	router.UpdateRouterInfo(edges)
}

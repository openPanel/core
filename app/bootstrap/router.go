package bootstrap

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/tools/convert"
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
	routerNodes := make([]router.Node, 1, len(nodes)+1)
	routerNodes[0] = self

	routerNodes = append(routerNodes, convert.NodesPbToRouter(nodes)...)

	router.SetNodes(routerNodes)

	return routerNodes
}

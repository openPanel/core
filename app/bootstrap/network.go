package bootstrap

import (
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/router"
)

func createEmptyNetGraph() {
	router.AddNodes([]router.Node{
		{
			Id: global.App.NodeInfo.ServerId,
			Address: router.Address{
				Ip:   global.App.NodeInfo.ServerIp,
				Port: global.App.NodeInfo.ServerPort,
			},
		},
	})
}

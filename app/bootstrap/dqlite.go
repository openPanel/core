package bootstrap

import (
	"context"

	"github.com/openPanel/core/app/config"
	. "github.com/openPanel/core/app/db/repo/shared"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/dqlite"
	"github.com/openPanel/core/app/manager/router"
)

// for initial node
func createDqlite() *shared.Client {
	return dqlite.CreateSharedDatabase(nil)
}


func initializeDqlite() error {
	err := config.SaveClusterInfo(global.App.ClusterInfo)
	if err != nil {
		return err
	}
	err = NodeRepo.AddNode(context.Background(),
		global.App.NodeInfo.ServerId,
		global.App.NodeInfo.ServerPublicIP.String(),
		global.App.NodeInfo.ServerPort,
	)
	if err != nil {
		return err
	}

	// store a cluster scoped token
	err = createToken()
	if err != nil {
		return err
	}
	return nil
}

// for success join, len(nodes) >= 1
func dqliteJoin(nodes []router.Node) *shared.Client {
	addrs := make([]string, len(nodes))
	for i, node := range nodes {
		addrs[i] = node.AddrPort.String()
	}
	return dqlite.CreateSharedDatabase(&addrs)
}

// at resume, len(addrs) >= 0, if len(addrs) == 0, it means that the node is the first node
func resumeDqlite(nodes []router.Node) *shared.Client {
	if len(nodes) == 0 {
		return dqlite.CreateSharedDatabase(nil)
	}

	addrs := make([]string, len(nodes))
	for i, node := range nodes {
		addrs[i] = node.AddrPort.String()
	}
	return dqlite.CreateSharedDatabase(&addrs)
}

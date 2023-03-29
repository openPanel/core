package rpc

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/dqlite"
)

var _ pb.DqliteConnectionServer = (*DqliteService)(nil)

type DqliteService struct{}

func (d *DqliteService) Connect(server pb.DqliteConnection_ConnectServer) error {
	conn := dqlite.NewServerRpcConn(server, "", global.App.NodeInfo.ServerId)
}

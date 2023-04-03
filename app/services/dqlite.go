package services

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/manager/dqlite"
)

var DqliteService pb.DqliteConnectionServer = new(dqliteService)

type dqliteService struct{}

func (d *dqliteService) Connect(server pb.DqliteConnection_ConnectServer) error {
	src, dst, err := getSrcAndDstFromContext(server.Context())
	if err != nil {
		return err
	}

	conn := dqlite.NewServerRpcConn(server, src, dst)

	dqlite.AcceptChan <- conn

	<-server.Context().Done()

	return server.Context().Err()
}

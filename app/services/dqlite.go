package services

import (
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/dqlite"
)

var DqliteService pb.DqliteConnectionServer = new(dqliteService)

type dqliteService struct{}

func (d *dqliteService) ServeDqlite(ss pb.DqliteConnection_ServeDqliteServer) error {
	src, dst, err := getSrcAndDstFromContext(ss.Context())
	if err != nil {
		log.Debugf("serve dqlite err: %v", err)
		return err
	}
	log.Debugf("serve dqlite: %v -> %v", src, dst)

	conn := dqlite.NewServerRpcConn(ss, src, dst)
	log.Debugf("dqlite conn created: %v -> %v", src, dst)

	dqlite.AcceptChan <- conn
	log.Debugf("dqlite conn insert: %v -> %v", src, dst)

	<-ss.Context().Done()
	log.Debugf("rpc conn closed: %v -> %v, %v", src, dst, ss.Context().Err())

	return conn.RpcConn.Context().Err()
}

package dqlite

import (
	"context"
	"net"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

func DialFunction(ctx context.Context, address string) (net.Conn, error) {
	if address == global.App.NodeInfo.ServerId {
		server, client := net.Pipe()
		AcceptChan <- client
		return server, nil
	}

	conn, err := rpcDialer.DialWithServerId(address)
	if err != nil {
		log.Debugf("dial dqlite err: %v", err)
		return nil, err
	}
	log.Debugf("dial dqlite: %v -> %v", global.App.NodeInfo.ServerId, address)

	client := pb.NewDqliteConnectionClient(conn)

	stream, err := client.ServeDqlite(ctx)
	if err != nil {
		log.Debugf("dial dqlite err: %v", err)
		return nil, err
	}
	log.Debugf("rpc conn created: %v -> %v", global.App.NodeInfo.ServerId, address)
	return NewClientRpcConn(stream, global.App.NodeInfo.ServerId, address), nil
}

var AcceptChan = make(chan net.Conn)

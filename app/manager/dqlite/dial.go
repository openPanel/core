package dqlite

import (
	"context"
	"net"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

// DialFunction ctx used to control first connect timeout, just ignore
func DialFunction(_ context.Context, address string) (net.Conn, error) {
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

	client := pb.NewDqliteConnectionClient(conn)

	stream, err := client.ServeDqlite(context.Background())
	if err != nil {
		log.Debugf("dial dqlite err: %v", err)
		return nil, err
	}
	return NewClientRpcConn(stream, global.App.NodeInfo.ServerId, address), nil
}

var AcceptChan = make(chan net.Conn)

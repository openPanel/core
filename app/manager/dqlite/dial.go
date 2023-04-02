package dqlite

import (
	"context"
	"net"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/rpc"
)

func DialFunction(ctx context.Context, address string) (net.Conn, error) {
	if address == global.App.NodeInfo.ServerId {
		server, client := net.Pipe()
		AcceptChan <- client
		return server, nil
	}

	conn, err := rpc.DialWithServerId(address)
	if err != nil {
		return nil, err
	}

	client := pb.NewDqliteConnectionClient(conn)

	stream, err := client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return NewClientRpcConn(stream, global.App.NodeInfo.ServerId, address), nil
}

var AcceptChan = make(chan net.Conn)

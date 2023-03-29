package dqlite

import (
	"net"

	"google.golang.org/grpc"
)

var _ net.Conn = (*ClientRpcConn)(nil)

type ClientRpcConn struct {
	rpcConn
}

func (c *ClientRpcConn) Close() error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	return c.stream.(grpc.ClientStream).CloseSend()
}

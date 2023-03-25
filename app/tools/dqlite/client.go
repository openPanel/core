package dqlite

import (
	"google.golang.org/grpc"
)

type ClientRpcConn struct {
	rpcConn
}

func (c *ClientRpcConn) Close() error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	return c.stream.(grpc.ClientStream).CloseSend()
}

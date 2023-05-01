package dqlite

import (
	"net"

	"google.golang.org/grpc"
)

var _ net.Conn = (*ClientRpcConn)(nil)

type ClientRpcConn struct {
	*RpcConn
}

func (c *ClientRpcConn) Close() error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	err := c.stream.(grpc.ClientStream).CloseSend()
	c.cancel()
	return err
}

func NewClientRpcConn(stream grpc.ClientStream, src, dst string) *ClientRpcConn {
	return &ClientRpcConn{
		RpcConn: NewRPCConn(
			NewRPCConnAddr(src),
			NewRPCConnAddr(dst),
			stream),
	}
}

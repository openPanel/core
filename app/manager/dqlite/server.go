package dqlite

import (
	"net"

	"google.golang.org/grpc"
)

var _ net.Conn = (*ServerRpcConn)(nil)

type ServerRpcConn struct {
	*RpcConn
}

// Close is a no-op. You can't close a gRPC stream on the server side.
func (s *ServerRpcConn) Close() error {
	s.cancel()
	return nil
}

func NewServerRpcConn(stream grpc.ServerStream, src, dst string) *ServerRpcConn {
	return &ServerRpcConn{
		RpcConn: NewRPCConn(
			NewRPCConnAddr(src),
			NewRPCConnAddr(dst),
			stream),
	}
}

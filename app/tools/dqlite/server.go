package dqlite

import (
	"google.golang.org/grpc"
)

type ServerRpcConn struct {
	rpcConn
}

// Close is a no-op. You can't close a gRPC stream on the server side.
func (s *ServerRpcConn) Close() error {
	return nil
}

func NewServerRpcConn(stream grpc.ServerStream, src string, dst string) *ServerRpcConn {
	return &ServerRpcConn{
		rpcConn{
			localAddr:  NewRPCConnAddr(src),
			remoteAddr: NewRPCConnAddr(dst),
			stream:     stream,
		},
	}
}

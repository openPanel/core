package dqlite

import (
	"net"
)

var _ net.Addr = (*RPCConnAddr)(nil)

type RPCConnAddr struct {
	name string
}

func NewRPCConnAddr(name string) RPCConnAddr {
	return RPCConnAddr{name: name}
}

func (R RPCConnAddr) Network() string {
	return "grpc conn"
}

func (R RPCConnAddr) String() string {
	return R.name
}

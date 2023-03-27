package dqlite

import (
	"context"
	"net"
)

func DialFunction(ctx context.Context, address string) (net.Conn, error) {
	panic("implement me") // TODO
}

var AcceptChan chan net.Conn

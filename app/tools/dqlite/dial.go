package dqlite

import (
	"context"
	"net"
)

func dqliteDialFunction(ctx context.Context, address string) (net.Conn, error) {
	panic("implement me") // TODO
}

var dqliteAcceptChan chan net.Conn

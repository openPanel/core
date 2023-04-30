package rpcDialer

import (
	"time"

	"github.com/zekroTJA/timedmap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"github.com/openPanel/core/app/constant"
)

var connCache = timedmap.New(time.Minute)

func getClientConnFromCache(remoteAddr string) *grpc.ClientConn {
	raw := connCache.GetValue(remoteAddr)
	if raw == nil {
		return nil
	}

	conn := raw.(*grpc.ClientConn)

	state := conn.GetState()
	if state == connectivity.TransientFailure || state == connectivity.Shutdown {
		connCache.Remove(remoteAddr)
		return nil
	}

	return conn
}

func connCleanCb(conn any) {
	_ = conn.(*grpc.ClientConn).Close()
}

func saveClientConnToCache(remoteAddr string, conn *grpc.ClientConn) {
	connCache.Set(remoteAddr, conn, constant.QuicClientConnCacheTimeout, connCleanCb)
}

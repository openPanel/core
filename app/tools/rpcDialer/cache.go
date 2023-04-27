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
	v := connCache.GetValue(remoteAddr).(*grpc.ClientConn)
	if v != nil {
		state := v.GetState()
		if state == connectivity.TransientFailure || state == connectivity.Shutdown {
			connCache.Remove(remoteAddr)
			return nil
		}
	}
	return v
}

func connCleanCb(conn any) {
	_ = conn.(*grpc.ClientConn).Close()
}

func saveClientConnToCache(remoteAddr string, conn *grpc.ClientConn) {
	connCache.Set(remoteAddr, conn, constant.QuicClientConnCacheTimeout, connCleanCb)
}

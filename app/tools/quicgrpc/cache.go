package quicgrpc

import (
	"context"
	"crypto/tls"

	"github.com/puzpuzpuz/xsync/v2"
	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/bootstrap/clean"
)

var connCache = xsync.NewMapOf[quic.Connection]()

func cacheableDial(ctx context.Context, tlsConf *tls.Config, addr string) (quic.Connection, error) {
	cachedConn, ok := connCache.Load(addr)
	if ok {
		return cachedConn, nil
	}
	conn, err := quic.DialAddrContext(ctx, addr, tlsConf, clientQuicConfig)
	if err != nil {
		return nil, err
	}

	connCache.Store(addr, conn)

	return conn, nil
}

func init() {
	clean.RegisterCleanup(func() {
		connCache.Range(func(key string, conn quic.Connection) bool {
			_ = conn.CloseWithError(0, "server closed")
			return true
		})
	})
}

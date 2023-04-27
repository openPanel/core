package constant

import (
	"time"

	"github.com/quic-go/quic-go"
)

var QuicConfig = &quic.Config{
	MaxIdleTimeout:  time.Minute,
	KeepAlivePeriod: time.Second * 30,
}

const QuicClientConnCacheTimeout = 10 * time.Minute
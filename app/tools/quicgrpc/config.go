package quicgrpc

import (
	"time"

	"github.com/quic-go/quic-go"
)

var clientQuicConfig = &quic.Config{
	MaxIdleTimeout:  time.Minute,
	KeepAlivePeriod: time.Second * 30,
	Tracer:          tracer,
}

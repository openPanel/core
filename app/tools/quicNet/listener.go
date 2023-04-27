package quicNet

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/constant"
)

var _ net.Listener = (*Listener)(nil)

type Listener struct {
	ql quic.EarlyListener
}

func Listen(ql quic.EarlyListener) net.Listener {
	return &Listener{ql}
}

func (l Listener) Accept() (net.Conn, error) {
	conn, err := l.ql.Accept(context.Background())
	if err != nil {
		return nil, err
	}

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}

	return &Conn{
		conn:   conn,
		stream: stream,
	}, nil
}

func (l Listener) Close() error {
	return l.ql.Close()
}

func (l Listener) Addr() net.Addr {
	return l.ql.Addr()
}

func NewQuicDialer(tlsConf *tls.Config) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, s string) (net.Conn, error) {
		conn, err := quic.DialAddrEarlyContext(ctx, s, tlsConf, constant.QuicConfig)
		if err != nil {
			return nil, err
		}
		return NewConn(conn)
	}
}

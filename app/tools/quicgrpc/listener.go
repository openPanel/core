package quicgrpc

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/global/log"
)

var _ net.Listener = (*Listener)(nil)

type Listener struct {
	ql        quic.Listener
	connQueue chan *Conn

	cancel context.CancelFunc
	ctx    context.Context
}

func Listen(ql quic.Listener) net.Listener {
	ctx, cancel := context.WithCancel(context.Background())

	listener := &Listener{
		ql,
		make(chan *Conn, 2),
		cancel,
		ctx,
	}

	go listener.acceptRoutine()

	return listener
}

func (l *Listener) acceptRoutine() {
	log.Debugf("quic listener accept routine started")

	for {
		select {
		case <-l.ctx.Done():
			return
		default:
			conn, err := l.ql.Accept(context.Background())
			if err != nil {
				continue
			}

			go l.acceptStreamRoutine(conn)
		}
	}
}

func (l *Listener) acceptStreamRoutine(conn quic.Connection) {
	for {
		select {
		case <-l.ctx.Done():
			return
		default:
			stream, err := conn.AcceptStream(context.Background())
			if err != nil {
				return
			}

			l.connQueue <- &Conn{
				localAddr:  conn.LocalAddr(),
				remoteAddr: conn.RemoteAddr(),
				stream:     stream,
			}
		}
	}
}

func (l *Listener) Accept() (net.Conn, error) {
	conn, ok := <-l.connQueue
	if !ok {
		return nil, net.ErrClosed
	}
	return conn, nil
}

func (l *Listener) Close() error {
	l.cancel()
	close(l.connQueue)
	return l.ql.Close()
}

func (l *Listener) Addr() net.Addr {
	return l.ql.Addr()
}

func NewQuicDialer(tlsConf *tls.Config) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, s string) (net.Conn, error) {
		conn, err := cacheableDial(ctx, tlsConf, s)
		if err != nil {
			return nil, err
		}

		return NewConn(conn)
	}
}

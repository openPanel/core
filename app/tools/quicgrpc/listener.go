package quicgrpc

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/constant"
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

	go listener.tryAccept()

	return listener
}

func (l *Listener) tryAccept() {
	log.Debugf("quic listener try accept loop start")

	for {
		select {
		case <-l.ctx.Done():
			log.Debugf("quic listener closed")
			return
		default:
			conn, err := l.ql.Accept(context.Background())
			if err != nil {
				log.Debugf("quic accept error: %v", err)
				continue
			}

			stream, err := conn.AcceptStream(context.Background())
			if err != nil {
				log.Debugf("quic accept stream error: %v", err)
				continue
			}

			l.connQueue <- &Conn{
				conn:   conn,
				stream: stream,
			}
		}
	}
}

func (l *Listener) Accept() (net.Conn, error) {
	conn, ok := <-l.connQueue
	if !ok {
		return nil, net.ErrClosed
	}
	log.Debugf("quic accept from conn queue")
	return conn, nil
}

func (l *Listener) Close() error {
	log.Debugf("quic listener closing")

	l.cancel()
	close(l.connQueue)
	return l.ql.Close()
}

func (l *Listener) Addr() net.Addr {
	return l.ql.Addr()
}

func NewQuicDialer(tlsConf *tls.Config) func(context.Context, string) (net.Conn, error) {
	// TODO: cache quic connection, only accept new stream if connection is alive
	return func(ctx context.Context, s string) (net.Conn, error) {
		conn, err := quic.DialAddrContext(ctx, s, tlsConf, constant.QuicConfig)
		if err != nil {
			return nil, err
		}
		log.Debugf("quic dial: %s", s)
		return NewConn(conn)
	}
}

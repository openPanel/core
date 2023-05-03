package quicgrpc

import (
	"context"
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

var _ net.Conn = (*Conn)(nil)

type Conn struct {
	localAddr  net.Addr
	remoteAddr net.Addr
	stream     quic.Stream
}

func NewConn(conn quic.Connection) (net.Conn, error) {
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	return &Conn{
		localAddr:  conn.LocalAddr(),
		remoteAddr: conn.RemoteAddr(),
		stream:     stream,
	}, nil
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.stream.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return c.stream.Write(b)
}

func (c *Conn) Close() error {
	return c.stream.Close()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *Conn) SetDeadline(t time.Time) error {
	return c.stream.SetDeadline(t)
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.stream.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.stream.SetWriteDeadline(t)
}

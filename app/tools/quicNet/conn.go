package quicNet

import (
	"context"
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

var _ net.Conn = (*Conn)(nil)

type Conn struct {
	conn   quic.Connection
	stream quic.Stream
}

func NewConn(conn quic.Connection) (net.Conn, error) {
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn:   conn,
		stream: stream,
	}, nil
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.stream.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return c.stream.Write(b)
}

func (c *Conn) Close() error {
	err := c.stream.Close()
	if err != nil {
		return err
	}
	return c.conn.CloseWithError(0, "quic connection closed")
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
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

package dqlite

import (
	"bufio"
	"context"
	"io"
	"net"
	"time"

	"github.com/openPanel/core/app/generated/pb"
)

// Stream grpc deprecate it, just use as a placeholder
type Stream interface {
	Context() context.Context
	SendMsg(m any) error
	RecvMsg(m any) error
}

var _ io.ReadWriter = (*RpcConn)(nil)

type RpcConn struct {
	localAddr  RPCConnAddr
	remoteAddr RPCConnAddr

	stream Stream

	readBuf *bufio.Reader

	ctx    context.Context
	cancel context.CancelFunc
}

func (r *RpcConn) Read(p []byte) (n int, err error) {
	n, err = r.readBuf.Read(p)
	return
}

func NewRPCConn(localAddr, remoteAddr RPCConnAddr, stream Stream) *RpcConn {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &RpcConn{
		localAddr:  localAddr,
		remoteAddr: remoteAddr,
		stream:     stream,
		ctx:        ctx,
		cancel:     cancel,
		readBuf: bufio.NewReader(
			&grpcReader{
				receive: stream.RecvMsg,
			},
		),
	}

	return conn
}

func (r *RpcConn) Context() context.Context {
	return r.ctx
}

type grpcReader struct {
	receive func(m any) error
}

func (r *grpcReader) Read(b []byte) (n int, err error) {
	resp := &pb.DqliteData{}
	err = r.receive(resp)
	if err != nil {
		return 0, err
	}

	return copy(b, resp.Data), nil
}

func (r *RpcConn) Write(b []byte) (n int, err error) {
	buf := make([]byte, len(b))
	copy(buf, b)

	err = r.stream.SendMsg(&pb.DqliteData{Data: buf})
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (r *RpcConn) LocalAddr() net.Addr {
	return r.localAddr
}

func (r *RpcConn) RemoteAddr() net.Addr {
	return r.remoteAddr
}

func (r *RpcConn) SetDeadline(t time.Time) error {
	return nil
}

func (r *RpcConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (r *RpcConn) SetWriteDeadline(t time.Time) error {
	return nil
}

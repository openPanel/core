package dqlite

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global/log"
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

	readLock  sync.Mutex
	writeLock sync.Mutex

	stream Stream

	ctx    context.Context
	cancel context.CancelFunc
}

func NewRPCConn(localAddr, remoteAddr RPCConnAddr, stream Stream) *RpcConn {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &RpcConn{
		localAddr:  localAddr,
		remoteAddr: remoteAddr,
		stream:     stream,
		ctx:        ctx,
		cancel:     cancel,
	}

	return conn
}

func (r *RpcConn) Context() context.Context {
	return r.ctx
}

func (r *RpcConn) Read(b []byte) (n int, err error) {
	log.Debugf("rpcConn.Read start")
	r.readLock.Lock()
	defer r.readLock.Unlock()

	resp := &pb.DqliteData{}
	err = r.stream.RecvMsg(resp)
	if err != nil {
		if err == io.EOF {
			r.cancel()
		}

		log.Debugf("rpcConn.Read: %v", err)
		return 0, err
	}
	log.Debugf("rpcConn.Read: %v", resp.Data)

	return copy(b, resp.Data), nil
}

func (r *RpcConn) Write(b []byte) (n int, err error) {
	log.Debugf("rpcConn.Write start")
	r.writeLock.Lock()
	defer r.writeLock.Unlock()

	buf := make([]byte, len(b))
	copy(buf, b)

	err = r.stream.SendMsg(&pb.DqliteData{Data: buf})
	if err != nil {
		log.Debugf("rpcConn.Write: %v", err)
		return 0, err
	}
	log.Debugf("rpcConn.Write: %v", b)

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

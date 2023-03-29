package dqlite

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/openPanel/core/app/generated/pb"
)

// Stream grpc deprecate it, just use as a placeholder
type Stream interface {
	Context() context.Context
	SendMsg(m any) error
	RecvMsg(m any) error
}

type rpcConn struct {
	localAddr  RPCConnAddr
	remoteAddr RPCConnAddr

	readLock  sync.Mutex
	writeLock sync.Mutex

	stream Stream
}

func (r *rpcConn) Read(b []byte) (n int, err error) {
	r.readLock.Lock()
	defer r.readLock.Unlock()

	resp := &pb.DqliteData{}
	err = r.stream.RecvMsg(resp)
	if err != nil {
		return 0, err
	}

	return copy(b, resp.Data), nil
}

func (r *rpcConn) Write(b []byte) (n int, err error) {
	r.writeLock.Lock()
	defer r.writeLock.Unlock()

	err = r.stream.SendMsg(&pb.DqliteData{Data: b})
	if err != nil {
		return 0, err
	}

	return len(b), nil
}

func (r *rpcConn) LocalAddr() net.Addr {
	return r.localAddr
}

func (r *rpcConn) RemoteAddr() net.Addr {
	return r.remoteAddr
}

func (r *rpcConn) SetDeadline(t time.Time) error {
	return nil
}

func (r *rpcConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (r *rpcConn) SetWriteDeadline(t time.Time) error {
	return nil
}

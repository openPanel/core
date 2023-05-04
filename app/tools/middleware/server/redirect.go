package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/rpcDialer"
)

func serverRedirectUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	dst, err := getDstFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if dst == global.App.NodeInfo.ServerId {
		return handler(ctx, req)
	}

	conn, err := rpcDialer.DialWithServerId(dst)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp = getReplyInstanceFromHandler(handler)
	err = conn.Invoke(ctx, info.FullMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func serverRedirectStreamInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Debugf("redirect interceptor called: %v", info.FullMethod)

	dst, err := getDstFromContext(ss.Context())
	if err != nil {
		log.Debugf("redirect err: %v", err)
		return err
	}

	if dst == global.App.NodeInfo.ServerId {
		log.Debugf("redirect dst is self")
		return handler(srv, ss)
	}
	log.Debugf("redirect dst: %v", dst)

	conn, err := rpcDialer.DialWithServerId(dst)
	if err != nil {
		log.Debugf("redirect err: %v", err)
		return err
	}

	desc := &grpc.StreamDesc{
		ClientStreams: info.IsClientStream,
		ServerStreams: info.IsServerStream,
	}

	stream, err := conn.NewStream(context.Background(), desc, info.FullMethod)
	if err != nil {
		log.Debugf("redirect err: %v", err)
		return err
	}
	log.Debugf("redirect stream created")

	return handler(srv, &redirectServerStream{stream})
}

var _ grpc.ServerStream = (*redirectServerStream)(nil)

// FIXME: implement headers and trailers correctly
type redirectServerStream struct {
	stream grpc.ClientStream
}

func (r *redirectServerStream) SetHeader(_ metadata.MD) error {
	return nil
}

func (r *redirectServerStream) SendHeader(_ metadata.MD) error {
	return nil
}

func (r *redirectServerStream) SetTrailer(_ metadata.MD) {
}

func (r *redirectServerStream) Context() context.Context {
	return r.stream.Context()
}

func (r *redirectServerStream) SendMsg(m interface{}) error {
	log.Debugf("SendMsg: %v", m)
	return r.stream.SendMsg(m)
}

func (r *redirectServerStream) RecvMsg(m interface{}) error {
	log.Debugf("RecvMsg: %v", m)
	return r.stream.RecvMsg(m)
}

var _ grpc.ServerStream = (*redirectServerStream)(nil)

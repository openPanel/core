package server

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/providers/zap/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/global"
)

var ServerUnaryInterceptorOption = grpc.ChainUnaryInterceptor(
	serverRedirectUnaryInterceptor,
	logging.UnaryServerInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc-server"))),
)

var ServerStreamInterceptorOption = grpc.ChainStreamInterceptor(
	serverRedirectStreamInterceptor,
	logging.StreamServerInterceptor(zap.InterceptorLogger(global.App.Logger.Named("grpc-server-stream"))),
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

	conn, err := rpc.DialWithServerId(dst)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp = getReplyInstanceFromHandler(handler)
	err = conn.Invoke(ctx, info.FullMethod, req, resp)
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
	dst, err := getDstFromContext(ss.Context())
	if err != nil {
		return err
	}

	if dst == global.App.NodeInfo.ServerId {
		return handler(srv, ss)
	}

	conn, err := rpc.DialWithServerId(dst)
	if err != nil {
		return err
	}

	desc := &grpc.StreamDesc{
		ClientStreams: info.IsClientStream,
		ServerStreams: info.IsServerStream,
	}

	stream, err := conn.NewStream(ss.Context(), desc, info.FullMethod)
	if err != nil {
		return err
	}

	if info.IsServerStream {
		go func() {
			for {
				m := getDataInstanceFromTransferFn(ss.RecvMsg)
				err := stream.RecvMsg(m)
				if err != nil {
					break
				}
				err = ss.SendMsg(m)
				if err != nil {
					break
				}
			}
		}()
	}

	if info.IsClientStream {
		go func() {
			for {
				m := getDataInstanceFromTransferFn(stream.RecvMsg)
				err := ss.RecvMsg(m)
				if err != nil {
					break
				}
				err = stream.SendMsg(m)
				if err != nil {
					break
				}
			}
		}()
	}

	<-ss.Context().Done()

	defer func(conn *grpc.ClientConn) {
		_ = stream.CloseSend()
		_ = conn.Close()
	}(conn)
	return nil
}

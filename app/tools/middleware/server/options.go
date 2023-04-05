package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/log"
)

func GetServerUnaryInterceptorOption() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		serverRedirectUnaryInterceptor,
		logging.UnaryServerInterceptor(log.InterceptorLogger(global.App.Logger.Named("grpc-server"))),
	)
}

func GetServerStreamInterceptorOption() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		serverRedirectStreamInterceptor,
		logging.StreamServerInterceptor(log.InterceptorLogger(global.App.Logger.Named("grpc-server-stream"))),
	)
}

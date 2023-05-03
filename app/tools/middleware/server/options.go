package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/tools/middleware/common"
)

func GetServerUnaryInterceptorOption() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		serverRedirectUnaryInterceptor,
		logging.UnaryServerInterceptor(common.InterceptorLogger(global.App.Logger.Named("grpc-server"))),
	)
}

func GetServerStreamInterceptorOption() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		serverRedirectStreamInterceptor,
		logging.StreamServerInterceptor(common.InterceptorLogger(global.App.Logger.Named("grpc-server-stream"))),
	)
}

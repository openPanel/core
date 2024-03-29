package services

import (
	"fmt"

	"github.com/quic-go/quic-go"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/detector/stop"
	"github.com/openPanel/core/app/tools/ca"
	"github.com/openPanel/core/app/tools/quicgrpc"
)

func StartRpcService() {
	grpcServer := newGrpcServer()

	tlsConfig, err := ca.GenerateRPCTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey, global.App.NodeInfo.ClusterCaCert)
	if err != nil {
		log.Panicf("error generating tls config: %v", err)
	}

	var listenAddr string

	if global.App.NodeInfo.IsIndirectIP {
		listenAddr = fmt.Sprintf("%s:%d", constant.DefaultListenIp, global.App.NodeInfo.ServerPort)
	} else {
		listenAddr = fmt.Sprintf("%s:%d", global.App.NodeInfo.ServerListenIP.String(), global.App.NodeInfo.ServerPort)
	}

	qle, err := quic.ListenAddr(listenAddr, tlsConfig, &constant.QuicConfig)
	if err != nil {
		log.Panicf("error listening: %v", err)
	}
	listener := quicgrpc.Listen(qle)

	stop.RegisterCleanup(func() {
		grpcServer.GracefulStop()
		log.Debug("grpc service stopped")
	}, constant.StopIDGRPCQUICServer, constant.StopIDLogger)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Panicf("error serving grpc: %v", err)
		}
	}()
}

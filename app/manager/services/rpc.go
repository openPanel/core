package services

import (
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quic-go/quic-go"
	"google.golang.org/grpc"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/middleware/server"
	"github.com/openPanel/core/app/tools/quicNet"
	"github.com/openPanel/core/app/tools/security"
)

var grpcServer = grpc.NewServer(server.ServerStreamInterceptorOption, server.ServerUnaryInterceptorOption)
var grpcMux = runtime.NewServeMux()

func StartRpcServiceBlocking() {
	tlsConfig, err := security.GenerateRPCTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey, global.App.ClusterInfo.CaCert)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	listenAddr := fmt.Sprintf("%s:%d", global.App.NodeInfo.ServerIp, global.App.NodeInfo.ServerPort)
	ql, err := quic.ListenAddr(listenAddr, tlsConfig, nil)
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	listener := quicNet.Listen(ql)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error serving grpc: %v", err)
	}
}

func StartHttpServiceBlocking() {
	unixListener, err := net.ListenUnix("unix", &net.UnixAddr{Name: constant.RpcUnixListenAddress, Net: "unix"})
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(unixListener); err != nil {
			log.Fatalf("error serving grpc: %v", err)
		}
	}()

	tlsConfig, err := security.GenerateHTTPTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	server := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", global.App.NodeInfo.ServerIp, global.App.NodeInfo.ServerPort),
		TLSConfig: tlsConfig,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("error serving http: %v", err)
	}
}

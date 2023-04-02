package services

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/third_party/OpenAPI"
)

func getGrpcMux() *runtime.ServeMux {
	grpcMux := runtime.NewServeMux()

	err := pb.RegisterInitializeServiceHandlerFromEndpoint(
		context.Background(),
		grpcMux,
		constant.RpcUnixListenAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("error registering grpc gateway: %v", err)
	}

	unixListener, err := net.Listen("unix", "")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(unixListener); err != nil {
			log.Fatalf("error serving loopback grpc: %v", err)
		}
	}()

	return grpcMux
}

func getServerHandler() http.HandlerFunc {
	router := httprouter.New()
	router.NotFound = getGrpcMux()
	router.Handler("GET", "/docs/*filepath", http.StripPrefix("/docs", OpenAPI.SwaggerUIHandler))
	return router.ServeHTTP
}

func StartHttpServiceBlocking() {
	tlsConfig, err := security.GenerateHTTPTLSConfig(global.App.NodeInfo.ServerCert, global.App.NodeInfo.ServerPrivateKey)
	if err != nil {
		log.Fatalf("error generating tls config: %v", err)
	}

	var addr string
	if global.App.NodeInfo.IsIndirectIP {
		addr = constant.DefaultListenIp.String()
	} else {
		addr = global.App.NodeInfo.ServerIp.String()
	}

	s := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", addr, global.App.NodeInfo.ServerPort),
		TLSConfig: tlsConfig,
		Handler:   getServerHandler(),
	}

	if err = s.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("error serving http: %v", err)
	}
}

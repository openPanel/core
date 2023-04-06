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

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/pb"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/middleware/gateway"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/third_party/OpenAPI"
)

func initGrpcGatewayMux() *runtime.ServeMux {
	unixListener, err := net.Listen("unix", "")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	go func() {
		grpcServer := newGrpcServer()

		clean.RegisterCleanup(func() {
			grpcServer.GracefulStop()
			log.Debug("unix grpc gateway stopped")
		})

		if err := grpcServer.Serve(unixListener); err != nil {
			log.Fatalf("error serving loop back grpc: %v", err)
		}
	}()

	grpcMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gateway.CustomMatcher),
	)

	err = pb.RegisterInitializeServiceHandlerFromEndpoint(
		context.Background(),
		grpcMux,
		unixListener.Addr().String(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, "unix", s)
			}),
		})
	if err != nil {
		log.Fatalf("error registering grpc gateway: %v", err)
	}

	return grpcMux
}

func wrapGrpcGatewayMux(mux *runtime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get(constant.RPCSourceMetadataKey) == "" {
			r.Header.Set(constant.RPCSourceMetadataKey, constant.RPCDefaultSource)
		}
		if r.Header.Get(constant.RPCDestinationMetadataKey) == "" {
			r.Header.Set(constant.RPCDestinationMetadataKey, global.App.NodeInfo.ServerId)
		}

		mux.ServeHTTP(w, r)
	})
}

func getServerHandler() http.HandlerFunc {
	router := httprouter.New()
	router.NotFound = wrapGrpcGatewayMux(initGrpcGatewayMux())
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
		addr = global.App.NodeInfo.ServerListenIP.String()
	}

	s := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", addr, global.App.NodeInfo.ServerPort),
		TLSConfig: tlsConfig,
		Handler:   getServerHandler(),
	}

	clean.RegisterCleanup(func() {
		err := s.Shutdown(context.Background())
		if err != nil {
			log.Warnf("error shutting down http server: %v", err)
		}
		log.Infof("outer grpc gateway stopped")
	})

	_ = s.ListenAndServeTLS("", "")
}

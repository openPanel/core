package bootstrap

import (
	"net"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/services"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/app/tools/utils"
)

// Start the first node of a cluster
func Start(listenIp net.IP, listenPort int) {
	requireInitialStartUp()

	commonInit()
	meta := generateNewNodeMeta(listenIp, listenPort)

	caCert, caKey, err := security.GenerateCACertificate()
	if err != nil {
		log.Fatalf("Failed to generate CA certificate: %v", err)
	}
	log.Info("CA certificate generated")

	localServerCert, err := security.SignCsr(caCert, caKey, meta.csr)
	if err != nil {
		log.Fatalf("Failed to sign local certificate: %v", err)
	}
	log.Info("Local certificate signed")

	node, err := saveNodeInfo(meta.serverId, meta.serverIp, meta.serverPort, localServerCert, meta.privateKey, meta.isIndirectIP)
	if err != nil {
		log.Fatalf("Failed to save node info: %v", err)
	}
	global.App.NodeInfo = node

	createEmptyNetGraph()

	global.App.ClusterInfo = global.ClusterInfo{
		CaCert: caCert,
		CaKey:  caKey,
	}

	global.App.DbShared = createDqlite()
	log.Info("Dqlite database configured")

	go services.StartRpcServiceBlocking()
	log.Infof("RPC service started on %s:%d", listenIp.String(), listenPort)

	go services.StartHttpServiceBlocking()
	log.Infof("HTTP service started on %s:%d", listenIp.String(), listenPort)

	e1 := config.SaveClusterInfo(global.App.ClusterInfo)

	_ = e1 // TODO: handle error

	utils.WaitExit()
}

// Join a cluster
func Join(listenIp net.IP, listenPort int, ip net.IP, port int, token string) {
	requireInitialStartUp()

	commonInit()

	meta := generateNewNodeMeta(listenIp, listenPort)

	_ = meta

}

// Resume resume a node to cluster
func Resume() {
	commonInit()
}

func commonInit() {
	initLogger()

	requireRoot()
	increaseUDPBufferSize()

	initLocalDatabase()
}

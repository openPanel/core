package bootstrap

import (
	"net"

	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/services"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/app/tools/utils"
)

// Start the first node of a cluster
func Start(listenIp net.IP, listenPort int) {
	commonInit()
	meta := generateNewNodeMeta(listenIp, listenPort)

	caCert, caKey, err := security.GenerateCACertificate()
	if err != nil {
		log.Fatalf("Failed to generate CA certificate: %v", err)
	}

	localServerCert, err := security.SignCsr(caCert, caKey, meta.Csr)
	if err != nil {
		log.Fatalf("Failed to sign local certificate: %v", err)
	}

	node, err := saveNodeInfo(meta.ServerId, meta.ServerIp, meta.ServerPort, localServerCert, meta.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to save node info: %v", err)
	}
	global.App.NodeInfo = node

	createEmptyNetGraph()

	global.App.ClusterInfo = global.ClusterInfo{
		CaCert: caCert,
		CaKey:  caKey,
	}

	createDqlite()

	go services.StartRpcServiceBlocking()
	go services.StartHttpServiceBlocking()

	utils.WaitExit()
}

// Join a cluster
func Join(listenIp net.IP, listenPort int, ip net.IP, port int, token string) {
	commonInit()
	generateNewNodeMeta(listenIp, listenPort)
}

// Resume resume a node to cluster
func Resume() {
	commonInit()
}

func commonInit() {
	initLogger()
	initLocalDatabase()
}

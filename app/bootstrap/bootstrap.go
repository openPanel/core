package bootstrap

import (
	"net"
	"net/netip"

	"github.com/openPanel/core/app/clients/http"
	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/cron"
	"github.com/openPanel/core/app/services"
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

	node := global.NodeInfo{
		ServerCert:       localServerCert,
		ServerPrivateKey: meta.privateKey,
		ServerId:         meta.serverId,
		ServerIp:         meta.serverIp,
		ServerPort:       meta.serverPort,
		ClusterCaCert:    caCert,
		IsIndirectIP:     meta.isIndirectIP,
	}
	err = config.SaveLocalNodeInfo(node)
	if err != nil {
		log.Fatalf("Failed to save node info: %v", err)
	}
	global.App.NodeInfo = node

	createEmptyNetGraph()

	global.App.ClusterInfo = global.ClusterInfo{
		CaKey: caKey,
	}

	global.App.DbShared = createDqlite()
	log.Info("Dqlite database configured")

	err = config.SaveClusterInfo(global.App.ClusterInfo)
	if err != nil {
		log.Fatalf("Failed to save cluster info: %v", err)
	}

	go services.StartRpcServiceBlocking()
	log.Infof("RPC service started on %s:%d", listenIp.String(), listenPort)

	go services.StartHttpServiceBlocking()
	log.Infof("HTTP service started on %s:%d", listenIp.String(), listenPort)

	// store a cluster scoped token
	err = createToken()
	if err != nil {
		log.Fatalf("Failed to create token: %v", err)
	}

	lateInit()

	utils.WaitExit()
}

// Join a cluster
func Join(listenIp net.IP, listenPort int, ip net.IP, port int, token string) {
	requireInitialStartUp()

	commonInit()

	meta := generateNewNodeMeta(listenIp, listenPort)
	target := netip.AddrPortFrom(netip.MustParseAddr(ip.String()), uint16(port))

	registerInfo, err := http.GetInitialInfo(target, meta.serverIp, meta.serverPort, token, meta.csr, meta.serverId)
	if err != nil {
		log.Fatalf("Failed to get initial info: %v", err)
	}

	node := global.NodeInfo{
		ServerId:         meta.serverId,
		ServerIp:         meta.serverIp,
		ServerPort:       meta.serverPort,
		ServerCert:       registerInfo.ClientCert,
		ServerPrivateKey: meta.privateKey,
		ClusterCaCert:    registerInfo.ClusterCACert,
		IsIndirectIP:     meta.isIndirectIP,
	}
	err = config.SaveLocalNodeInfo(node)
	if err != nil {
		log.Fatalf("Failed to save node info: %v", err)
	}

	createFullNetGraphAtJoin(registerInfo)

	global.App.DbShared = joinDqlite(registerInfo)
	global.App.ClusterInfo, err = config.LoadClusterInfo()
	if err != nil {
		log.Fatalf("Failed to load cluster info: %v", err)
	}

	go services.StartRpcServiceBlocking()
	log.Infof("RPC service started on %s:%d", listenIp.String(), listenPort)

	go services.StartHttpServiceBlocking()
	log.Infof("HTTP service started on %s:%d", listenIp.String(), listenPort)

	lateInit()

	utils.WaitExit()
}

// Resume resume a node to cluster
func Resume() {
	commonInit()

	var err error

	global.App.NodeInfo, err = config.LoadLocalNodeInfo()
	if err != nil {
		log.Fatalf("Failed to load node info: %v", err)
	}

	// Check if the current node was the only node in the cluster
	// the last time it exited the cluster.
	// If not, should try to get the current cluster status from one of the neighbors,
	// otherwise start a single node.
	cache, err := config.LoadNodesCache()
	if err != nil {
		log.Fatalf("Failed to load nodes cache: %v", err)
	}
	if len(cache) > 0 {
		targets := make([]rpc.Target, len(cache))
		for i, node := range cache {
			targets[i] = rpc.Target{
				ServerId: node.Id,
				AddrPort: node.AddrPort,
			}
		}

		addrs, err := rpc.UpdateRouterInfo(targets)
		if err != nil {
			log.Fatalf("Failed to update router info: %v", err)
		}

		global.App.DbShared = resumeDqlite(addrs)
	} else {
		global.App.DbShared = resumeDqlite(nil)
	}

	global.App.ClusterInfo, err = config.LoadClusterInfo()
	if err != nil {
		log.Fatalf("Failed to load cluster info: %v", err)
	}

	go services.StartRpcServiceBlocking()
	log.Infof("RPC service started on %s:%d", global.App.NodeInfo.ServerIp, global.App.NodeInfo.ServerPort)

	go services.StartHttpServiceBlocking()
	log.Infof("HTTP service started on %s:%d", global.App.NodeInfo.ServerIp, global.App.NodeInfo.ServerPort)

	lateInit()
}

func commonInit() {
	initLogger()

	requireRoot()
	increaseUDPBufferSize()

	initLocalDatabase()
}

func lateInit() {
	go cron.Start()
}

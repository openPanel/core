package bootstrap

import (
	"net"

	"github.com/openPanel/core/app/bootstrap/clean"
	"github.com/openPanel/core/app/clients/http"
	"github.com/openPanel/core/app/clients/rpc"
	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/manager/cron"
	"github.com/openPanel/core/app/manager/router"
	"github.com/openPanel/core/app/services"
	"github.com/openPanel/core/app/tools/ca"
	"github.com/openPanel/core/app/tools/convert"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

// Create the first node of a cluster
func Create(listenIp net.IP, listenPort int) {
	requireFirstStartUp()
	commonInit()

	meta := generateNewNode(listenIp, listenPort)

	caCert, caKey, err := ca.GenerateCACertificate()
	if err != nil {
		log.Panicf("Failed to generate CA certificate: %v", err)
	}
	log.Info("CA certificate generated")

	localServerCert, err := ca.SignCsr(caCert, caKey, meta.csr)
	if err != nil {
		log.Panicf("Failed to sign local certificate: %v", err)
	}
	log.Info("Local certificate signed")

	node := global.NodeInfo{
		ServerCert:       localServerCert,
		ServerPrivateKey: meta.privateKey,
		ServerId:         meta.serverId,
		ServerPublicIP:   meta.serverPublicIp,
		ServerListenIP:   meta.serverListenIp,
		ServerPort:       meta.serverPort,
		ClusterCaCert:    caCert,
		IsIndirectIP:     meta.isIndirectIP,
	}

	err = config.SaveLocalNodeInfo(node)
	if err != nil {
		log.Panicf("Failed to save node info: %v", err)
	}
	log.Infof("Node info saved")

	global.App.NodeInfo = node

	createEmptyNetGraph()

	global.App.ClusterInfo = global.ClusterInfo{
		CaKey: caKey,
	}

	global.App.DbShared = createDqlite()
	log.Info("Dqlite database configured")

	err = initializeDqlite()
	if err != nil {
		log.Panicf("Failed to initialize database in dqlite: %v", err)
	}

	startServices()

	clean.RunEndless()
}

// Join a cluster
func Join(listenIp net.IP, listenPort int, ip net.IP, port int, token string) {
	requireFirstStartUp()
	commonInit()

	initialized := false
	defer func() {
		if !initialized {
			err := cleanData()
			if err != nil {
				panic(err)
			}
		}
	}()

	meta := generateNewNode(listenIp, listenPort)
	target := netUtils.NewAddrPortWithIP(ip, port)

	// contains known node in the cluster
	initialInfo, err := http.GetClusterInfo(target, token)
	if err != nil {
		log.Errorf("Failed to get initial info: %v", err)
		return
	}

	routerNodes := loadAndSaveInitialNodes(initialInfo.Nodes, router.Node{
		Id:       meta.serverId,
		AddrPort: netUtils.NewAddrPortWithIP(meta.serverPublicIp, meta.serverPort),
	})

	// no need to test latency with itself
	linkStates := router.EstimateLatencies(routerNodes[1:], meta.serverId)
	log.Infof("Latencies estimated")

	shutdownTempSvr, err := bootTempServer(listenIp, listenPort)
	if err != nil {
		log.Infof("Failed to boot temp server: %v", err)
		return
	}

	registerInfo, err := http.RegisterNewNode(
		target,
		meta.serverPublicIp,
		meta.serverPort,
		token,
		meta.csr,
		meta.serverId,
		convert.LinkStatesRouterToPb(linkStates))
	if err != nil {
		log.Panicf("Failed to register: %v", err)
	}

	shutdownTempSvr()

	global.App.NodeInfo = global.NodeInfo{
		ServerId:         meta.serverId,
		ServerPublicIP:   meta.serverPublicIp,
		ServerListenIP:   meta.serverListenIp,
		ServerPort:       meta.serverPort,
		ServerPrivateKey: meta.privateKey,
		IsIndirectIP:     meta.isIndirectIP,
		ServerCert:       registerInfo.ClientCert,
		ClusterCaCert:    registerInfo.ClusterCACert,
	}
	err = config.SaveLocalNodeInfo(global.App.NodeInfo)
	if err != nil {
		log.Panicf("Failed to save node info: %v", err)
	}

	router.MergeLinkStates(convert.LinkStatesPbToRouter(registerInfo.LinkStates))

	global.App.DbShared = dqliteJoin(routerNodes[1:])
	global.App.ClusterInfo, err = config.LoadClusterInfo()
	if err != nil {
		log.Panicf("Failed to load cluster info: %v", err)
	}

	startServices()

	initialized = true
	clean.RunEndless()
}

// Resume resume a node to cluster
func Resume() {
	requireNonFirstStartUp()
	commonInit()

	var err error

	global.App.NodeInfo, err = config.LoadLocalNodeInfo()
	if err != nil {
		log.Panicf("Failed to load node info: %v", err)
	}

	// Check if the current node was the only node in the cluster
	// the last time it exited the cluster.
	// If not, should try to get the current cluster status from one of the neighbors,
	// otherwise start a single node.
	cacheNodes, err := config.LoadNodesCache()
	if err != nil {
		log.Panicf("Failed to load nodes cache: %v", err)
	}

	// There exists other nodes in the cluster when the node exited last time.
	if len(cacheNodes) > 1 {
		targets := make([]rpc.Target, 0, len(cacheNodes)-1)
		for _, node := range cacheNodes {
			if node.Id == global.App.NodeInfo.ServerId {
				continue
			}
			targets = append(targets, rpc.Target{
				ServerId: node.Id,
				AddrPort: node.AddrPort,
			})
		}

		shutdownTempSvr, err := bootTempServer(global.App.NodeInfo.ServerListenIP, global.App.NodeInfo.ServerPort)
		nodes, lst, err := rpc.SyncNodesAndLinkStates(targets)
		if err != nil {
			log.Panicf("Failed to update router info: %v", err)
		}
		shutdownTempSvr()

		router.SetNodes(nodes)
		router.MergeLinkStates(lst)

		global.App.DbShared = resumeDqlite(nodes)
	} else {
		global.App.DbShared = resumeDqlite(nil)
	}

	global.App.ClusterInfo, err = config.LoadClusterInfo()
	if err != nil {
		log.Panicf("Failed to load cluster info: %v", err)
	}

	startServices()

	clean.RunEndless()
}

func commonInit() {
	initLogger()

	requireRoot()
	requireEnoughUDPBuffer()

	initLocalDatabase()
}

func startServices() {
	go cron.Start()
	log.Infof("Cron service started")

	go services.StartRpcServiceBlocking()
	log.Infof("RPC service started on %s:%d", global.App.NodeInfo.ServerListenIP, global.App.NodeInfo.ServerPort)

	go services.StartHttpServiceBlocking()
	log.Infof("HTTP service started on %s:%d", global.App.NodeInfo.ServerListenIP, global.App.NodeInfo.ServerPort)
}

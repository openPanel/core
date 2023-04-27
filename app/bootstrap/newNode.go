package bootstrap

import (
	"net"

	"github.com/google/uuid"

	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

type newNodeMeta struct {
	serverId       string
	serverPublicIp net.IP
	serverListenIp net.IP
	serverPort     int
	csr            []byte
	privateKey     []byte

	isIndirectIP bool
}

func generateNewNodeMeta(ip net.IP, port int) newNodeMeta {
	serverId := uuid.New().String()

	var publicIp net.IP
	var indirect bool
	var err error

	if ip.IsUnspecified() || ip.IsPrivate() || !ip.IsGlobalUnicast() {
		var serverIps []net.IP
		serverIps, indirect, err = netUtils.GetPublicIP()
		if err != nil {
			log.Panicf("Failed to get public IP: %v", err)
		}

		if len(serverIps) > 1 {
			log.Infof("Multiple public IPs found: %v", serverIps)
			log.Infof("Using first IP: %v", serverIps[0])
		}
		publicIp = serverIps[0]
	} else {
		publicIp = ip
	}

	serverPort := port

	signingCsr, privateKey, err := security.GenerateCertificateSigningRequest(serverId)
	if err != nil {
		log.Panicf("Failed to generate certificate signing request: %v", err)
	}

	return newNodeMeta{
		serverId:       serverId,
		serverPublicIp: publicIp,
		serverListenIp: ip,
		serverPort:     serverPort,
		csr:            signingCsr,
		isIndirectIP:   indirect,
		privateKey:     privateKey,
	}
}

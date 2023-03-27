package bootstrap

import (
	"net"

	"github.com/google/uuid"

	"github.com/openPanel/core/app/global/log"
	"github.com/openPanel/core/app/tools/security"
	"github.com/openPanel/core/app/tools/utils/netUtils"
)

type newNodeMeta struct {
	ServerId   string
	ServerIp   net.IP
	ServerPort int
	Csr        []byte
	PrivateKey []byte
}

func generateNewNodeMeta(ip net.IP, port int) newNodeMeta {
	serverId := uuid.New().String()

	var serverIp net.IP

	if ip.IsUnspecified() {
		serverIps, err := netUtils.GetPublicIP()
		if err != nil {
			log.Fatalf("Failed to get public IP: %v", err)
		}

		if len(serverIps) > 1 {
			log.Warnf("Multiple public IPs found: %v", serverIps)
			log.Warnf("Using first IP: %v", serverIps[0])
		}
		serverIp = serverIps[0]
	} else {
		serverIp = ip
	}

	serverPort := port

	signingCsr, privateKey, err := security.GenerateCertificateSigningRequest(serverId)
	if err != nil {
		log.Fatalf("Failed to generate certificate signing request: %v", err)
	}

	return newNodeMeta{
		ServerId:   serverId,
		ServerIp:   serverIp,
		ServerPort: serverPort,
		Csr:        signingCsr,
		PrivateKey: privateKey,
	}
}

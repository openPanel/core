package bootstrap

import (
	slog "log"
	"net"
	"os"
	"os/user"
	"strconv"

	"github.com/lorenzosaino/go-sysctl"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"

	"github.com/openPanel/core/app/global/log"
)

func requireRoot() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if u.Uid != "0" {
		log.Fatal("This program requires root privileges")
	}
}

func saveNodeInfo(id string, ip net.IP, port int, cert, key []byte, indirect bool) (global.NodeInfo, error) {
	i := global.NodeInfo{
		ServerCert:       cert,
		ServerId:         id,
		ServerIp:         ip,
		ServerPort:       port,
		ServerPrivateKey: key,
		IsIndirectIP:     indirect,
	}
	err := config.SaveLocalNodeInfo(i)
	return i, err
}

func loadNodeInfo() (id string, ip net.IP, port int, cert, key []byte) {
	i, err := config.LoadLocalNodeInfo()
	if err != nil {
		log.Fatalf("Failed to load node info: %v", err)
	}
	return i.ServerId, i.ServerIp, i.ServerPort, i.ServerCert, i.ServerPrivateKey
}

func requireInitialStartUp() {
	// check file exist
	if _, err := os.Stat(constant.DefaultDataDir + string(os.PathSeparator) + constant.DefaultLocalSqliteFilename); err == nil {
		slog.Fatal("This instance has already been initialized, boot with no command to resume the server")
	} else if !os.IsNotExist(err) {
		slog.Fatalf("Failed to check if database file exists: %v", err)
	}
}

func increaseUDPBufferSize() {
	v, err := sysctl.Get(constant.SysctlUdpBufferSizeKey)
	if err != nil {
		log.Fatalf("Failed to read udp buffer: %v", err)
	}
	vNum, err := strconv.Atoi(v)
	if vNum < constant.SysctlUdpBufferSizeValue {
		err := sysctl.Set(constant.SysctlUdpBufferSizeKey, strconv.Itoa(constant.SysctlUdpBufferSizeValue))
		if err != nil {
			log.Fatalf("Failed to increase udp buffer: %v", err)
		}
		log.Infof("Increased UDP buffer size to %d", constant.SysctlUdpBufferSizeValue)
	} else {
		log.Infof("UDP buffer size is %d, not need to increase", vNum)
	}
}

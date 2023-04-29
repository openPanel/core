package bootstrap

import (
	slog "log"
	"os"
	"os/user"
	"strconv"

	"github.com/lorenzosaino/go-sysctl"

	"github.com/openPanel/core/app/constant"
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

func requireFirstStartUp() {
	// check file exist
	if _, err := os.Stat(constant.DefaultDataDir + string(os.PathSeparator) + constant.DefaultLocalSqliteFilename); err == nil {
		slog.Fatal("This instance has already been initialized, boot with no command to resume the server")
	} else if !os.IsNotExist(err) {
		slog.Fatalf("Failed to check if database file exists: %v", err)
	}
}

func requireNonFirstStartUp() {
	// check file exist
	if _, err := os.Stat(constant.DefaultDataDir + string(os.PathSeparator) + constant.DefaultLocalSqliteFilename); err != nil {
		slog.Fatalf("Failed to check if database file exists: %v", err)
	}
}

func cleanData() error {
	// delete everything in data dir if init failed
	err := os.RemoveAll(constant.DefaultDataDir)
	if err != nil {
		slog.Fatalf("Failed to clean data dir: %v", err)
	}
	slog.Println("Initialization failed, data dir cleaned")
	return nil
}

func requireEnoughUDPBuffer() {
	v, err := sysctl.Get(constant.SysctlUdpBufferSizeKey)
	if err != nil {
		log.Fatalf("Failed to read udp buffer: %v", err)
	}
	vNum, err := strconv.Atoi(v)
	if vNum < constant.SysctlUdpBufferSizeValue {
		err = sysctl.Set(constant.SysctlUdpBufferSizeKey, strconv.Itoa(constant.SysctlUdpBufferSizeValue))
		if err != nil {
			log.Fatalf("Failed to increase udp buffer: %v", err)
		}
		log.Infof("Increased UDP buffer size to %d", constant.SysctlUdpBufferSizeValue)
	} else {
		log.Infof("UDP buffer size is %d, not need to increase", vNum)
	}
}

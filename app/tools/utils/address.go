package utils

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/openPanel/core/app/constant"
)

func AssertPublicAddress(address string) (net.IP, int) {
	parts := strings.Split(address, ":")
	if len(parts) == 0 || len(parts) > 2 {
		//fatalWithMessage("Invalid address " + address)
		log.Println()
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		panic("Invalid IP address " + parts[0])
	}

	if !ip.IsGlobalUnicast() {
		panic("IP address is not global unicast address " + ip.String())
	}
	var port int
	var err error

	if len(parts) == 2 {
		port, err = strconv.Atoi(parts[1])

		if err != nil {
			panic("Invalid port " + parts[1])
		}

		if port < 1 || port > 65535 {
			panic("Invalid port " + parts[1])
		}
	} else {
		port = constant.DefaultListenPort
	}

	return ip, port
}

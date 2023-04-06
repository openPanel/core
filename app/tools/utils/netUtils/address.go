package netUtils

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
)

func CheckPublicIp(ip net.IP) error {
	if global.IsDev() {
		return nil
	}

	if ip.IsUnspecified() {
		return nil
	}
	if !ip.IsGlobalUnicast() {
		return errors.New("IP address is not global unicast address " + ip.String())
	}
	if ip.IsPrivate() {
		return errors.New("IP address is private " + ip.String())
	}
	return nil
}

func AssertPublicAddress(address string) (net.IP, int) {
	parts := strings.Split(address, ":")
	if len(parts) == 0 || len(parts) > 2 {
		log.Println("Invalid address " + address)
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		log.Fatal("Invalid IP address " + parts[0])
	}

	err := CheckPublicIp(ip)
	if err != nil {
		log.Fatal(err)
	}

	var port int

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

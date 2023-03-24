package netUtils

import (
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/openPanel/core/app/global/log"
)

func GetPublicIP() (ips []net.IP, err error) {
	localIPs, err := getPublicIPsFromLocal()
	if err != nil {
		log.Error("Failed to get public IP from local network interfaces", zap.Error(err))
		return nil, err
	}
	if len(localIPs) > 0 {
		return localIPs, nil
	} else {
		log.Debug("No public IP found from local network interfaces")
		cfip, err := getPublicIPWithCloudflare()
		if err == nil {
			return []net.IP{cfip}, nil
		}
		log.Error("Failed to get public IP from Cloudflare", zap.Error(err))
		ipifyip, err := getPublicIPWithIpify()
		if err == nil {
			return []net.IP{ipifyip}, nil
		}
		log.Fatal("Failed to get public IP from ipify", zap.Error(err))
		return nil, errors.New("failed to get public IP")
	}
}

func getPublicIPsFromLocal() (ips []net.IP, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.IsGlobalUnicast() && !v.IP.IsPrivate() {
					ips = append(ips, v.IP)
				}
			case *net.IPAddr:
				if v.IP.IsGlobalUnicast() && !v.IP.IsPrivate() {
					ips = append(ips, v.IP)
				}
			}
		}
	}
	return ips, nil
}

func getPublicIPWithCloudflare() (net.IP, error) {
	const CloudflareTraceEndpoint = "https://dash.cloudflare.com/cdn-cgi/trace"

	resp, err := http.Get(CloudflareTraceEndpoint)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respTextStr := string(respText)

	lines := strings.Split(respTextStr, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ip=") {
			ip := net.ParseIP(strings.TrimPrefix(line, "ip="))
			if ip == nil {
				return nil, errors.New("could not parse ip")
			}
			log.Warnf("found ip: %s from cloudflare trace, may not be correct", ip.String())
			return ip, nil
		}
	}
	return nil, errors.New("could not find ip in response")
}

func getPublicIPWithIpify() (net.IP, error) {
	const IpifyEndpoint = "https://api.ipify.org"

	resp, err := http.Get(IpifyEndpoint)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respTextStr := string(respText)

	ip := net.ParseIP(respTextStr)
	if ip == nil {
		return nil, errors.New("could not parse ip")
	}
	log.Warnf("found ip: %s from ipify, may not be correct", ip.String())
	return ip, nil
}

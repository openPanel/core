package netUtils

import (
	"net"
	"net/netip"
)

func NewAddrPortWithString(ip string, port int) netip.AddrPort {
	return netip.AddrPortFrom(netip.MustParseAddr(ip), uint16(port))
}

func NewAddrPortWithIP(ip net.IP, port int) netip.AddrPort {
	return netip.AddrPortFrom(netip.MustParseAddr(ip.String()), uint16(port))
}

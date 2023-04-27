package cmd

import (
	"net"
	"strconv"

	"github.com/urfave/cli/v2"
)

var _ cli.Generic = (*IP)(nil)
var _ cli.Generic = (*Port)(nil)

type IP net.IP

func NewIP(ip net.IP) *IP {
	return (*IP)(&ip)
}

func (ip *IP) Set(s string) error {
	parsed := net.ParseIP(s)
	if parsed == nil {
		return net.InvalidAddrError(s)
	}
	*ip = IP(parsed)
	return nil
}

func (ip *IP) String() string {
	return (*net.IP)(ip).String()
}

type Port int

func NewPort(port int) *Port {
	return (*Port)(&port)
}

func (p *Port) Set(s string) error {
	port, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if port < 1 || port > 65535 {
		return net.InvalidAddrError(s)
	}
	*p = Port(port)
	return nil
}

func (p *Port) String() string {
	return strconv.Itoa(int(*p))
}

package http

import (
	"net"
	"net/netip"
	"time"
)

func TcpPing(target netip.AddrPort) (int, error) {
	dialer := net.Dialer{
		KeepAlive: -1,
		Timeout:   5 * time.Second,
	}

	tcpStart := time.Now()
	tcpConn, err := dialer.Dial("tcp", target.String())
	if err != nil {
		return 0, err
	}
	defer func(tcpConn net.Conn) {
		_ = tcpConn.Close()
	}(tcpConn)
	tcpLatency := time.Since(tcpStart).Milliseconds()
	return int(tcpLatency), nil
}

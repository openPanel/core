package bootstrap

import (
	"context"
	"net"
	"strconv"
)

// boot temporary server to handle the latency check
func bootTempServer(listenIp net.IP, listenPort int) (func(), error) {
	cleanups := make([]func(), 0)

	isOk := false
	clean := func() {
		for _, cleanup := range cleanups {
			cleanup()
		}
	}

	defer func() {
		if !isOk {
			clean()
		}
	}()

	// inform the handler to stop
	ctx, cancel := context.WithCancel(context.Background())
	cleanups = append(cleanups, cancel)

	// boot Tcp Server
	tcpAddr, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort(listenIp.String(), strconv.Itoa(listenPort)))
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
					return
				default:
				}
				continue
			}
			go func(conn net.Conn) {
				_ = conn.Close()
			}(conn)
		}
	}()
	cleanups = append(cleanups, func() {
		_ = tcpListener.Close()
	})

	isOk = true
	return clean, nil
}

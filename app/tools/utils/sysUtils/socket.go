package sysUtils

import (
	"net"
	"os"
	"syscall"
)

func SocketPair() (net.Conn, net.Conn, error) {
	fds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}

	c1, err := fdToFileConn(fds[0])
	if err != nil {
		return nil, nil, err
	}

	c2, err := fdToFileConn(fds[1])
	if err != nil {
		_ = c1.Close()
		return nil, nil, err
	}

	return c1, c2, err
}

func fdToFileConn(fd int) (net.Conn, error) {
	f := os.NewFile(uintptr(fd), "")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	return net.FileConn(f)
}

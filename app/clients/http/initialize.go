package http

import (
	"net"
)

func GetInitialInfo(ip net.IP, port int, token string) {
	type InitialRequest struct {
		Ip       net.IP `json:"ip"`
		Port     int    `json:"port"`
		ServerID string `json:"serverID"`
		Token    string `json:"token"`
		Csr      []byte `json:"csr"`
	}
}

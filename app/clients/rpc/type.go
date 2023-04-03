package rpc

import (
	"net/netip"
)

type Target struct {
	ServerId string
	AddrPort netip.AddrPort
}

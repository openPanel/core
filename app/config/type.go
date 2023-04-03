package config

import (
	"net/netip"
)

type NodeCacheEntry struct {
	Id       string
	AddrPort netip.AddrPort
}

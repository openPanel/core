package router

import (
	"sync"
)

// RouterInfo The map store the latency between two nodes
var routerInfo = make(map[string]map[string]int32)
var routerInfoLock = sync.RWMutex{}

// RouterDecision The map store the decision of the router, value define the next hop
var routerDecision = make(map[string]string)

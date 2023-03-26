package router

// implement a default route algorithm
// It directly connects to other node if it is possible
// For unreachable node, it will use the lowest latency node to connect to it
// This algorithm will be used before received all node link state.

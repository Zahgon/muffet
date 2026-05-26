package main

import "sync"

type hostThrottlerPool struct {
	requestPerSecond, maxConnectionsPerHost int
	hostMap                                 sync.Map
}

func newHostThrottlerPool(requestPerSecond, maxConnectionsPerHost int) *hostThrottlerPool {
	_ = "STUB: not implemented"
	return nil
}

func (p *hostThrottlerPool) Get(name string) *hostThrottler { _ = "STUB: not implemented"; return nil }

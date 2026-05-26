package main

import "go.uber.org/ratelimit"

type hostThrottler struct {
	limiter     ratelimit.Limiter
	connections semaphore
}

func newHostThrottler(requestPerSecond, maxConnectionsPerHost int) *hostThrottler {
	_ = "STUB: not implemented"
	return nil
}

func (t *hostThrottler) Request() { _ = "STUB: not implemented"; return }

func (t *hostThrottler) Release() { _ = "STUB: not implemented"; return }

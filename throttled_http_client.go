package main

import (
	"net/http"
	"net/url"
)

type throttledHttpClient struct {
	client            httpClient
	connections       semaphore
	hostThrottlerPool *hostThrottlerPool
}

func newThrottledHttpClient(c httpClient, requestPerSecond int, maxConnections, maxConnectionsPerHost int) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}

func (c *throttledHttpClient) Get(u *url.URL, header http.Header) (httpResponse, error) {
	_ = "STUB: not implemented"
	return *new(httpResponse), nil
}

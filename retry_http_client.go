package main

import (
	"net/http"
	"net/url"
	"time"
)

type retryHttpClient struct {
	client       httpClient
	maxCount     uint
	initialDelay time.Duration
}

func newRetryHttpClient(c httpClient, maxCount uint, initialDelay time.Duration) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}

func (c *retryHttpClient) Get(u *url.URL, header http.Header) (httpResponse, error) {
	_ = "STUB: not implemented"
	return *new(httpResponse), nil
}

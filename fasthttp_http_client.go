package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type fasthttpHttpClient struct {
	client  *fasthttp.Client
	timeout time.Duration
	header  http.Header
}

func newFasthttpHttpClient(c *fasthttp.Client, timeout time.Duration, header http.Header) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}

func (c *fasthttpHttpClient) Get(u *url.URL, header http.Header) (httpResponse, error) {
	_ = "STUB: not implemented"
	return *new(httpResponse), nil
}

// Some HTTP servers require "Accept" headers set explicitly.

func includeHeader(h http.Header, k string) bool { _ = "STUB: not implemented"; return false }

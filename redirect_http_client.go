package main

import (
	"net/http"
	"net/url"
)

type redirectHttpClient struct {
	client          httpClient
	maxRedirections int
}

func newRedirectHttpClient(c httpClient, maxRedirections int) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}

func (c *redirectHttpClient) Get(u *url.URL, header http.Header) (httpResponse, error) {
	_ = "STUB: not implemented"
	return *new(httpResponse), nil
}

func parseCookies(s string) []*http.Cookie { _ = "STUB: not implemented"; return nil }

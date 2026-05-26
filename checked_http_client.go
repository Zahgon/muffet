package main

import (
	"net/http"
	"net/url"
)

type checkedHttpClient struct {
	client              httpClient
	acceptedStatusCodes statusCodeSet
}

func newCheckedHttpClient(c httpClient, acceptedStatusCodes statusCodeSet) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}

func (c *checkedHttpClient) Get(u *url.URL, header http.Header) (httpResponse, error) {
	_ = "STUB: not implemented"
	return *new(httpResponse), nil
}

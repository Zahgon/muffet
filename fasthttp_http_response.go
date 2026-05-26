package main

import (
	"github.com/valyala/fasthttp"
)

type fasthttpHttpResponse struct {
	url      *fasthttp.URI
	response *fasthttp.Response
}

func newFasthttpHttpResponse(u *fasthttp.URI, r *fasthttp.Response) httpResponse {
	_ = "STUB: not implemented"
	return *new(httpResponse)
}

func (r fasthttpHttpResponse) URL() string { _ = "STUB: not implemented"; return "" }

func (r fasthttpHttpResponse) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (r fasthttpHttpResponse) Header(key string) string { _ = "STUB: not implemented"; return "" }

func (r fasthttpHttpResponse) Body() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

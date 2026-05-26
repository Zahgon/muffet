package main

import (
	"net/url"
)

type htmlPage struct {
	url       *url.URL
	fragments map[string]struct{}
	links     map[string]error
}

func newHtmlPage(u *url.URL, fragments map[string]struct{}, links map[string]error) *htmlPage {
	_ = "STUB: not implemented"
	return nil
}

func (p *htmlPage) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (p *htmlPage) Fragments() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (p *htmlPage) Links() map[string]error { _ = "STUB: not implemented"; return nil }

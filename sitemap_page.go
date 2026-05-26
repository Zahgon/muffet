package main

import (
	"net/url"
)

type sitemapPage struct {
	url   *url.URL
	links map[string]error
}

func newSitemapPage(u *url.URL, links map[string]error) *sitemapPage {
	_ = "STUB: not implemented"
	return nil
}

func (p *sitemapPage) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (p *sitemapPage) Fragments() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (p *sitemapPage) Links() map[string]error { _ = "STUB: not implemented"; return nil }

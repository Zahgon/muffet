package main

import (
	"net/url"
)

type sitemapPageParser struct {
	linkFilterer linkFilterer
}

func newSitemapPageParser(f linkFilterer) *sitemapPageParser { _ = "STUB: not implemented"; return nil }

func (p *sitemapPageParser) Parse(u *url.URL, typ string, bs []byte) (page, error) {
	_ = "STUB: not implemented"
	return *new(page), nil
}

// TODO Detect XML files as sitemaps.

// TODO Detect XML files as sitemap indices.

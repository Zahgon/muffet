package main

import (
	"net/url"
)

type sitemapFetcher struct {
	client httpClient
}

func newSitemapFetcher(c httpClient) *sitemapFetcher { _ = "STUB: not implemented"; return nil }

func (f *sitemapFetcher) Fetch(uu *url.URL) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*sitemapFetcher) formatGetError(err error) error { _ = "STUB: not implemented"; return nil }

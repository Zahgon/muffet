package main

import (
	"net/url"

	"github.com/temoto/robotstxt"
)

type robotsTxtFetcher struct {
	client httpClient
}

func newRobotsTxtFetcher(c httpClient) *robotsTxtFetcher { _ = "STUB: not implemented"; return nil }

func (f *robotsTxtFetcher) Fetch(uu *url.URL) (*robotstxt.RobotsData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*robotsTxtFetcher) formatError(err error) error { _ = "STUB: not implemented"; return nil }

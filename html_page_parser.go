package main

import (
	"net/url"
)

type htmlPageParser struct {
	linkFinder linkFinder
}

func newHtmlPageParser(f linkFinder) *htmlPageParser { _ = "STUB: not implemented"; return nil }

func (p htmlPageParser) Parse(u *url.URL, typ string, body []byte) (page, error) {
	_ = "STUB: not implemented"
	return *new(page), nil
}

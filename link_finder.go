package main

import (
	"net/url"
	"regexp"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var atomToAttributes = map[atom.Atom][]string{
	atom.A:      {"href"},
	atom.Frame:  {"src"},
	atom.Iframe: {"src"},
	atom.Img:    {"src"},
	atom.Link:   {"href"},
	atom.Script: {"src"},
	atom.Source: {"src", "srcset"},
	atom.Track:  {"src"},
	atom.Meta:   {"content"},
}

var imageDescriptorPattern = regexp.MustCompile(`(\S)\s+\S+\s*$`)

type linkFinder struct {
	linkFilterer linkFilterer
}

func newLinkFinder(f linkFilterer) linkFinder { _ = "STUB: not implemented"; return *new(linkFinder) }

func (f linkFinder) Find(n *html.Node, base *url.URL) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// `preconnect` and `dns-prefetch` links are not HTTP resources.

func (f linkFinder) parseLinks(n *html.Node, a string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (linkFinder) trimUrl(s string) string { _ = "STUB: not implemented"; return "" }

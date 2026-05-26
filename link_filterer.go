package main

import (
	"net/url"
	"regexp"
)

var validSchemes = map[string]struct{}{
	"":      {},
	"http":  {},
	"https": {},
}

type linkFilterer struct {
	excludedPatterns []*regexp.Regexp
	includedPatterns []*regexp.Regexp
}

func newLinkFilterer(es []*regexp.Regexp, is []*regexp.Regexp) linkFilterer {
	_ = "STUB: not implemented"
	return *new(linkFilterer)
}

func (f linkFilterer) IsValid(u *url.URL) bool { _ = "STUB: not implemented"; return false }

func (f linkFilterer) isLinkExcluded(u string) bool { _ = "STUB: not implemented"; return false }

func (f linkFilterer) isLinkIncluded(u string) bool { _ = "STUB: not implemented"; return false }

func (f linkFilterer) matches(u string, rs []*regexp.Regexp) bool {
	_ = "STUB: not implemented"
	return false
}

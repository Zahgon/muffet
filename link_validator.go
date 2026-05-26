package main

import (
	"net/url"

	"github.com/temoto/robotstxt"
)

type linkValidator struct {
	hostname    string
	sitemapURLs map[string]struct{}
	robotsData  *robotstxt.RobotsData
}

func newLinkValidator(hostname string, robotsData *robotstxt.RobotsData, sitemap map[string]struct{}) *linkValidator {
	_ = "STUB: not implemented"
	return nil
}

// Validate validates a link and returns true if it is valid as one of an HTML page.
func (v *linkValidator) Validate(u *url.URL) bool { _ = "STUB: not implemented"; return false }

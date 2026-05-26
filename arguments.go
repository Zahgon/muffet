package main

import (
	"net/http"
	"regexp"
)

type arguments struct {
	RawAcceptedStatusCodes string   `long:"accepted-status-codes" value-name:"<codes>" default:"200..300" description:"Accepted HTTP response status codes (e.g. '200..300,403')"`
	BufferSize             int      `short:"b" long:"buffer-size" value-name:"<size>" default:"4096" description:"HTTP response buffer size in bytes"`
	MaxConnections         int      `short:"c" long:"max-connections" value-name:"<count>" default:"512" description:"Maximum number of HTTP connections"`
	MaxConnectionsPerHost  int      `long:"max-connections-per-host" value-name:"<count>" default:"512" description:"Maximum number of HTTP connections per host"`
	MaxResponseBodySize    int      `long:"max-response-body-size" value-name:"<size>" default:"10000000" description:"Maximum response body size to read"`
	RawExcludedPatterns    []string `short:"e" long:"exclude" value-name:"<pattern>..." description:"Exclude URLs matched with given regular expressions"`
	RawIncludedPatterns    []string `short:"i" long:"include" value-name:"<pattern>..." description:"Include URLs matched with given regular expressions"`
	FollowRobotsTxt        bool     `long:"follow-robots-txt" description:"Follow robots.txt when scraping pages"`
	FollowSitemapXML       bool     `long:"follow-sitemap-xml" description:"Scrape only pages listed in sitemap.xml (deprecated)"`
	RawHeaders             []string `long:"header" value-name:"<header>..." description:"Custom headers"`
	// TODO Remove a short option.
	IgnoreFragments bool   `short:"f" long:"ignore-fragments" description:"Ignore URL fragments"`
	MaxRetries      uint   `long:"max-retries" value-name:"<count>" default:"0" description:"Maximum retry count for network errors"`
	DnsResolver     string `long:"dns-resolver" value-name:"<address>" description:"Custom DNS resolver"`
	Format          string `long:"format" description:"Output format" default:"text" choice:"text" choice:"json" choice:"junit"`
	// TODO Remove this option.
	JSONOutput bool `long:"json" description:"Output results in JSON (deprecated)"`
	// TODO Remove this option.
	VerboseJSON bool `long:"experimental-verbose-json" description:"Include successful results in JSON (deprecated)"`
	// TODO Remove this option.
	JUnitOutput         bool   `long:"junit" description:"Output results as JUnit XML file (deprecated)"`
	MaxRedirections     int    `short:"r" long:"max-redirections" value-name:"<count>" default:"64" description:"Maximum number of redirections"`
	RateLimit           int    `long:"rate-limit" value-name:"<rate>" description:"Max requests per second"`
	Timeout             int    `short:"t" long:"timeout" value-name:"<seconds>" default:"10" description:"Timeout for HTTP requests in seconds"`
	Verbose             bool   `short:"v" long:"verbose" description:"Show successful results too"`
	Proxy               string `long:"proxy" value-name:"<host>" description:"HTTP proxy host"`
	SkipTLSVerification bool   `long:"skip-tls-verification" description:"Skip TLS certificate verification"`
	OnePageOnly         bool   `long:"one-page-only" description:"Only check links found in the given URL"`
	Color               color  `long:"color" description:"Color output" choice:"auto" choice:"always" choice:"never" default:"auto"`
	Help                bool   `short:"h" long:"help" description:"Show this help"`
	Version             bool   `long:"version" description:"Show version"`
	URL                 string
	AcceptedStatusCodes statusCodeSet
	ExcludedPatterns    []*regexp.Regexp
	IncludePatterns     []*regexp.Regexp
	Header              http.Header
}

func getArguments(ss []string) (*arguments, error) { _ = "STUB: not implemented"; return nil, nil }

func help() string { _ = "STUB: not implemented"; return "" }

// Parse() is run here to show default values in help.
// This seems to be a bug in go-flags.
// nolint:errcheck

func compileRegexps(regexps []string) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseHeaders(headers []string) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}

func reconcileDeprecatedArguments(args *arguments) { _ = "STUB: not implemented"; return }

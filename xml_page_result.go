package main

type xmlPageResult struct {
	Url      string `xml:"name,attr"`
	Total    int    `xml:"tests,attr"`
	Failures int    `xml:"failures,attr"`
	Skipped  int    `xml:"skipped,attr"`
	// spell-checker: disable-next-line
	Links []*xmlLinkResult `xml:"testcase"`
}

type xmlLinkResult struct {
	Url string `xml:"name,attr"`
	// spell-checker: disable-next-line
	Source  string          `xml:"classname,attr"`
	Failure *xmlLinkFailure `xml:"failure"`
}

type xmlLinkFailure struct {
	Message string `xml:"message,attr"`
}

func newXMLPageResult(pr *pageResult) *xmlPageResult { _ = "STUB: not implemented"; return nil }

// TODO: Consider adding information skipped links, if that can be tracked.

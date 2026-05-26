package main

type linkFetcher struct {
	client      httpClient
	pageParsers []pageParser
	cache       cache
	options     linkFetcherOptions
}

type fetchResult struct {
	StatusCode int
	Page       page
}

func newLinkFetcher(c httpClient, ps []pageParser, o linkFetcherOptions) *linkFetcher {
	_ = "STUB: not implemented"
	return nil
}

// Fetch fetches a link and returns a successful status code and optionally HTML page, or an error.
func (f *linkFetcher) Fetch(u string) (int, page, error) {
	_ = "STUB: not implemented"
	return 0, *new(page), nil
}

// TODO Support text fragments.

func (f *linkFetcher) sendRequestWithCache(u string) (int, page, error) {
	_ = "STUB: not implemented"
	return 0, *new(page), nil
}

func (f *linkFetcher) sendRequest(s string) (int, page, error) {
	_ = "STUB: not implemented"
	return 0, *new(page), nil
}

func separateFragment(s string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

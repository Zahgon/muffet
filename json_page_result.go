package main

type jsonPageResult struct {
	URL   string `json:"url"`
	Links []any  `json:"links"`
}

type jsonSuccessLinkResult struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
}

type jsonErrorLinkResult struct {
	URL   string `json:"url"`
	Error string `json:"error"`
}

func newJSONPageResult(r *pageResult, verbose bool) *jsonPageResult {
	_ = "STUB: not implemented"
	return nil
}

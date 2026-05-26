package main

type pageChecker struct {
	fetcher       *linkFetcher
	linkValidator *linkValidator
	daemonManager *daemonManager
	results       chan *pageResult
	donePages     concurrentStringSet
	onePageOnly   bool
}

func newPageChecker(f *linkFetcher, v *linkValidator, onePageOnly bool) *pageChecker {
	_ = "STUB: not implemented"
	return nil
}

func (c *pageChecker) Results() <-chan *pageResult { _ = "STUB: not implemented"; return nil }

func (c *pageChecker) Check(page page) { _ = "STUB: not implemented"; return }

func (c *pageChecker) checkPage(p page) { _ = "STUB: not implemented"; return }

func (c *pageChecker) addPage(p page) { _ = "STUB: not implemented"; return }

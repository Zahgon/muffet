package main

import "sync"

type concurrentStringSet struct {
	set *sync.Map
}

func newConcurrentStringSet() concurrentStringSet {
	_ = "STUB: not implemented"
	return *new(concurrentStringSet)
}

func (c concurrentStringSet) Add(s string) bool { _ = "STUB: not implemented"; return false }

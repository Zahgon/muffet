package main

import "sync"

type cache struct {
	locks  *sync.Map
	values *sync.Map
}

func newCache() cache { _ = "STUB: not implemented"; return *new(cache) }

func (c cache) LoadOrStore(key string) (any, func(any)) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

package main

type statusCodeSet map[statusCodeRange]struct{}

func parseStatusCodeSet(value string) (statusCodeSet, error) {
	_ = "STUB: not implemented"
	return *new(statusCodeSet), nil
}

func (s statusCodeSet) Contains(code int) bool { _ = "STUB: not implemented"; return false }

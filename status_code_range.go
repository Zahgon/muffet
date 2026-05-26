package main

type statusCodeRange struct {
	start int
	end   int
}

func parseStatusCodeRange(s string) (*statusCodeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r statusCodeRange) Contains(code int) bool { _ = "STUB: not implemented"; return false }

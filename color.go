package main

type color string

const (
	auto   color = "auto"
	always color = "always"
	never  color = "never"
)

func isColorEnabled(c color, terminal bool) bool { _ = "STUB: not implemented"; return false }

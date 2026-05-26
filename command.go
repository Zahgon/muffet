package main

import (
	"io"
)

type command struct {
	stdout, stderr    io.Writer
	terminal          bool
	httpClientFactory httpClientFactory
}

func newCommand(stdout, stderr io.Writer, terminal bool, f httpClientFactory) *command {
	_ = "STUB: not implemented"
	return nil
}

func (c *command) Run(args []string) bool { _ = "STUB: not implemented"; return false }

func (c *command) runWithError(ss []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *command) printResultsInJSON(rc <-chan *pageResult, verbose bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *command) printResultsInJUnitXML(rc <-chan *pageResult) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// spell-checker: disable-next-line

// spell-checker: disable-next-line

func (c *command) print(xs ...any) { _ = "STUB: not implemented"; return }

func (c *command) printError(xs ...any) { _ = "STUB: not implemented"; return }

// Do not check --color option here because this can be used on argument parsing errors.

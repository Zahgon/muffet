package main

import (
	"github.com/logrusorgru/aurora/v3"
)

type pageResultFormatter struct {
	verbose bool
	aurora  aurora.Aurora
}

func newPageResultFormatter(verbose bool, color bool) *pageResultFormatter {
	_ = "STUB: not implemented"
	return nil
}

func (f *pageResultFormatter) Format(r *pageResult) string { _ = "STUB: not implemented"; return "" }

func (f *pageResultFormatter) formatSuccessLinkResults(rs []*successLinkResult) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *pageResultFormatter) formatErrorLinkResults(rs []*errorLinkResult) []string {
	_ = "STUB: not implemented"
	return nil
}

func formatMessages(ss []string) []string { _ = "STUB: not implemented"; return nil }

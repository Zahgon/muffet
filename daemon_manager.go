package main

import "sync"

type daemonManager struct {
	daemons   chan func()
	waitGroup *sync.WaitGroup
}

func newDaemonManager(capacity int) *daemonManager { _ = "STUB: not implemented"; return nil }

func (m daemonManager) Add(f func()) { _ = "STUB: not implemented"; return }

func (m daemonManager) Run() { _ = "STUB: not implemented"; return }

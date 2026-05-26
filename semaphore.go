package main

type semaphore struct {
	channel chan bool
}

func newSemaphore(n int) semaphore { _ = "STUB: not implemented"; return *new(semaphore) }

func (s semaphore) Request() { _ = "STUB: not implemented"; return }

func (s semaphore) Release() { _ = "STUB: not implemented"; return }

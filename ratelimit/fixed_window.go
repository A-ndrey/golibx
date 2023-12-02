package ratelimit

import (
	"sync"
	"time"
)

type FixedWindow struct {
	window time.Duration
	limit  uint

	counter uint
	start   time.Time
	mx      sync.Mutex
}

func NewFixedWindow(limit uint, window time.Duration) *FixedWindow {
	fw := FixedWindow{
		window: window,
		limit:  limit,
	}

	return &fw
}

func (fw *FixedWindow) Inc() bool {
	fw.mx.Lock()
	defer fw.mx.Unlock()

	diff := time.Now().Sub(fw.start)
	if diff > fw.window {
		fw.start = time.Now()
		fw.counter = 0
	}

	if fw.counter >= fw.limit {
		return false
	}

	fw.counter++

	return true
}

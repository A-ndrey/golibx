package ratelimit

import (
	"testing"
	"time"
)

func TestFixedWindow(t *testing.T) {
	fw := NewFixedWindow(10, time.Minute)

	var okCnt, failCnt int
	for i := 0; i < 100; i++ {
		if fw.Inc() {
			okCnt++
		} else {
			failCnt++
		}
	}

	if okCnt != 10 || failCnt != 90 {
		t.Errorf("FixedWindow{limit:%d, window:%s}, ok:%d, fail:%d", fw.limit, fw.window, okCnt, failCnt)
	}
}

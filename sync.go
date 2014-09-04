package logr

import "sync"

type Sync struct {
	Out Logger

	mu sync.Mutex
}

func (l *Sync) Log(e Entry) {
	l.mu.Lock()
	l.Out.Log(e)
	l.mu.Unlock()
}

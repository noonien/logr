package logr

import "sync"

type Sync struct {
	Out Logger

	mu sync.Mutex
}

func (l *Sync) Log(e Entry) {
	if l.Out == nil {
		return
	}

	l.mu.Lock()
	l.Out.Log(e)
	l.mu.Unlock()
}

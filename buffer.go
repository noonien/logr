package logr

import (
	"sync"
	"time"
)

type Buffer struct {
	Out      Logger
	Entries  int
	Interval time.Duration

	mu sync.Mutex
	e  []Entry
	qc chan struct{}
}

func (l *Buffer) Log(e Entry) {
	if l.Entries == 0 && l.Interval == 0 {
		panic("logr: Buffer.Entries and Buffer.Interval cannot both be 0")
	}

	// Protect e, qc
	l.mu.Lock()
	defer l.mu.Unlock()

	// Append the entry
	l.e = append(l.e, e)
	elen := len(l.e)

	// If we've reached out maximum entry count, forward them all
	if elen == l.Entries {
		l.forward()
		return
	}

	// This is the first entry in the list, start a goroutine to send it,
	// along with the others after the maximum interval has passed
	if elen == 1 && l.Interval > 0 {
		l.qc = make(chan struct{}, 1)
		go func() {
			select {
			// Forward entries after interval has passed
			case <-time.After(l.Interval):
				l.mu.Lock()
				defer l.mu.Unlock()
				l.forward()

			// Return if signaled to
			case <-l.qc:
			}
			l.qc = nil
		}()
	}
}

func (l *Buffer) forward() {
	// Signal the waiting goroutie to quit since we're sending all the existing
	// entries
	if l.qc != nil {
		l.qc <- struct{}{}
	}

	// Just return if there's nothing to forward to
	if l.Out == nil {
		l.e = nil
		return
	}

	// Forward entries
	for _, e := range l.e {
		l.Out.Log(e)
	}

	// Clear entries
	l.e = nil
}

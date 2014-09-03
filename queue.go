package logr

type Queue struct {
	entries []Entry
}

func (l *Queue) Log(e Entry) {
	if e == nil {
		return
	}
	l.entries = append(l.entries, e)
}

func (l *Queue) LogBatch(e []Entry) {
	for _, entry := range e {
		if entry == nil {
			continue
		}
		l.entries = append(l.entries, entry)
	}
}

func (l *Queue) Len() int {
	return len(l.entries)
}

func (l *Queue) Peek() Entry {
	if len(l.entries) == 0 {
		return nil
	}
	return l.entries[0]
}

func (l *Queue) Pop() Entry {
	if len(l.entries) == 0 {
		return nil
	}
	e := l.entries[0]
	l.entries = l.entries[1:]
	return e
}

func (l *Queue) Clear() {
	l.entries = nil
}

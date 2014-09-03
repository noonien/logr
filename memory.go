package logr

type Memory struct {
	entries []Entry
}

func (l *Memory) Log(e Entry) {
	if e == nil {
		return
	}
	l.entries = append(l.entries, e)
}

func (l *Memory) Len() int {
	return len(l.entries)
}

func (l *Memory) Peek() Entry {
	if len(l.entries) == 0 {
		return nil
	}
	return l.entries[0]
}

func (l *Memory) Pop() Entry {
	if len(l.entries) == 0 {
		return nil
	}
	e := l.entries[0]
	l.entries = l.entries[1:]
	return e
}

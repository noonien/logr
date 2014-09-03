package logr

type Entry map[string]interface{}

type Logger interface {
	Log(Entry)
}

type BatchLogger interface {
	LogBatch([]Entry)
}

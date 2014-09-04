package logr

type Discard struct{}

func (_ Discard) Log(e Entry) {}

func (_ Discard) LogBatch(e []Entry) {}

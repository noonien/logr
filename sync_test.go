package logr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSync(t *testing.T) {
	Convey("Given a Sync logger", t, func() {
		q := Queue{}
		log := &Sync{
			Out: &q,
		}

		Convey("It can be used as a Logger", func() {
			_, ok := interface{}(log).(Logger)
			So(ok, ShouldEqual, true)
		})

		Convey("It doesn't panic when an output is not assigned", func() {
			log.Out = nil
			So(func() { log.Log(nil) }, ShouldNotPanic)
		})

		Convey("It forwards the enty to the output", func() {
			entry := Entry{"test": "test"}
			log.Log(entry)
			So(q.Peek(), ShouldEqual, entry)
		})
	})
}

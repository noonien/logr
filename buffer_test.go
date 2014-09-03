package logr

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuffer(t *testing.T) {
	Convey("Given a Buffer logger", t, func() {
		q := Queue{}
		log := &Buffer{
			Out: &q,
		}

		Convey("It can be used as a Logger", func() {
			_, ok := interface{}(log).(Logger)
			So(ok, ShouldEqual, true)
		})

		Convey("It should panic when Entries and Interval are not set", func() {
			So(func() { log.Log(nil) }, ShouldPanic)
		})

		Convey("It doesn't panic when an output is not assigned", func() {
			log.Out = nil
			log.Entries = 1
			So(func() { log.Log(nil) }, ShouldNotPanic)
		})

		Convey("It should forward entries only when maximum count has been reached", func() {
			log.Entries = 2

			log.Log(Entry{"test": "test"})
			So(q.Len(), ShouldEqual, 0)

			log.Log(Entry{"test": "test"})
			So(q.Len(), ShouldEqual, 2)
		})

		Convey("It should forward entries only when maximum interval has been reached", func() {
			log.Interval = 10 * time.Millisecond

			log.Log(Entry{"test": "test"})
			So(q.Len(), ShouldEqual, 0)

			time.Sleep(2 * log.Interval)
			So(q.Len(), ShouldEqual, 1)
		})

		Convey("It should forward entries after either maximum interval or entry count has been reached", func() {
			log.Entries = 2
			log.Interval = 10 * time.Millisecond

			log.Log(Entry{"test": "test"})
			So(q.Len(), ShouldEqual, 0)

			log.Log(Entry{"test": "test"})
			So(q.Len(), ShouldEqual, 2)

			q.Clear()

			time.Sleep(2 * log.Interval)
			So(q.Len(), ShouldEqual, 0)

			log.Log(Entry{"test": "test"})
			time.Sleep(2 * log.Interval)
			So(q.Len(), ShouldEqual, 1)
		})
	})
}

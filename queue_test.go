package logr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	Convey("Given a Queue logger", t, func() {
		log := &Queue{}

		Convey("It can be used as a Logger", func() {
			_, ok := interface{}(log).(Logger)
			So(ok, ShouldEqual, true)
		})

		Convey("It can be used as a BatchLogger", func() {
			_, ok := interface{}(log).(BatchLogger)
			So(ok, ShouldEqual, true)
		})

		Convey("Adding a nil entry is ignored", func() {
			log.Log(nil)
			So(log.Len(), ShouldEqual, 0)

			log.LogBatch([]Entry{nil})
			So(log.Len(), ShouldEqual, 0)
		})

		Convey("When empty, return nil Entry", func() {
			So(log.Pop(), ShouldBeNil)
			So(log.Peek(), ShouldBeNil)
		})

		Convey("When adding entries", func() {
			entry1 := Entry{"test1": "test1"}
			entry2 := Entry{"test2": "test2"}

			log.LogBatch([]Entry{entry1})
			log.Log(entry2)

			Convey("The Length is increased", func() {
				So(log.Len(), ShouldEqual, 2)
			})

			Convey("The first entry should be returned by Peek", func() {
				So(log.Peek(), ShouldEqual, entry1)
			})

			Convey("The first entry should be returned and removed by Pop", func() {
				So(log.Pop(), ShouldEqual, entry1)
				So(log.Peek(), ShouldEqual, entry2)
			})

			Convey("The entries should be removed by Clear", func() {
				log.Clear()
				So(log.Len(), ShouldEqual, 0)
			})

			Convey("Entries are returned in a FIFO order", func() {
				So(log.Pop(), ShouldEqual, entry1)
				So(log.Pop(), ShouldEqual, entry2)
			})
		})
	})
}

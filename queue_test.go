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

		Convey("Adding a nil entry is ignored", func() {
			log.Log(nil)
			So(log.Len(), ShouldEqual, 0)
		})

		Convey("When empty, return nil Entry", func() {
			So(log.Pop(), ShouldBeNil)
			So(log.Peek(), ShouldBeNil)
		})

		Convey("When adding an entry", func() {
			entry := Entry{"test": "testing"}
			log.Log(entry)

			Convey("The Length is increased", func() {
				So(log.Len(), ShouldEqual, 1)
			})

			Convey("The entry should be returned by Peek", func() {
				So(log.Peek(), ShouldEqual, entry)
			})

			Convey("The entry should be returned and removed by Pop", func() {
				So(log.Pop(), ShouldEqual, entry)
				So(log.Peek(), ShouldBeNil)
			})

			Convey("Entries are returned in a FIFO order", func() {
				entry2 := Entry{"test2": "test2"}
				log.Log(entry2)
				So(log.Pop(), ShouldEqual, entry)
				So(log.Pop(), ShouldEqual, entry2)
			})
		})
	})
}

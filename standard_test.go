package logr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStandard(t *testing.T) {
	Convey("Given a Standard logger", t, func() {
		q := Queue{}
		log := &Standard{
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

		Convey("Calling its Print* methods adds an entry to the output logger", func() {
			log.Print("Print")
			log.Printf("Printf")
			log.Println("Println")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"message": "Print"})
			So(q.Pop(), ShouldResemble, Entry{"message": "Printf"})
			So(q.Pop(), ShouldResemble, Entry{"message": "Println"})

		})

		Convey("Calling its Panic* methods adds an entry to the output logger and panics", func() {
			So(func() { log.Panic("Panic") }, ShouldPanicWith, "Panic")
			So(func() { log.Panicf("Panicf") }, ShouldPanicWith, "Panicf")
			So(func() { log.Panicln("Panicln") }, ShouldPanicWith, "Panicln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"message": "Panic"})
			So(q.Pop(), ShouldResemble, Entry{"message": "Panicf"})
			So(q.Pop(), ShouldResemble, Entry{"message": "Panicln"})
		})
	})
}

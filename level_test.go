package logr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLevel(t *testing.T) {
	Convey("Given a Level logger", t, func() {
		q := Queue{}
		log := &Level{
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

		Convey("Calling its Debug* methods adds an entry to the output logger", func() {
			log.Debug("Debug")
			log.Debugf("Debugf")
			log.Debugln("Debugln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"level": "debug", "message": "Debug"})
			So(q.Pop(), ShouldResemble, Entry{"level": "debug", "message": "Debugf"})
			So(q.Pop(), ShouldResemble, Entry{"level": "debug", "message": "Debugln"})
		})
		Convey("Calling its Info* methods adds an entry to the output logger", func() {
			log.Info("Info")
			log.Infof("Infof")
			log.Infoln("Infoln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"level": "info", "message": "Info"})
			So(q.Pop(), ShouldResemble, Entry{"level": "info", "message": "Infof"})
			So(q.Pop(), ShouldResemble, Entry{"level": "info", "message": "Infoln"})
		})

		Convey("Calling its Warn* methods adds an entry to the output logger", func() {
			log.Warn("Warn")
			log.Warnf("Warnf")
			log.Warnln("Warnln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"level": "warn", "message": "Warn"})
			So(q.Pop(), ShouldResemble, Entry{"level": "warn", "message": "Warnf"})
			So(q.Pop(), ShouldResemble, Entry{"level": "warn", "message": "Warnln"})
		})

		Convey("Calling its Error* methods adds an entry to the output logger", func() {
			log.Error("Error")
			log.Errorf("Errorf")
			log.Errorln("Errorln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"level": "error", "message": "Error"})
			So(q.Pop(), ShouldResemble, Entry{"level": "error", "message": "Errorf"})
			So(q.Pop(), ShouldResemble, Entry{"level": "error", "message": "Errorln"})
		})

		Convey("Calling its Panic* methods adds an entry to the output logger and panics", func() {
			So(func() { log.Panic("Panic") }, ShouldPanicWith, "Panic")
			So(func() { log.Panicf("Panicf") }, ShouldPanicWith, "Panicf")
			So(func() { log.Panicln("Panicln") }, ShouldPanicWith, "Panicln")

			So(q.Len(), ShouldEqual, 3)
			So(q.Pop(), ShouldResemble, Entry{"level": "panic", "message": "Panic"})
			So(q.Pop(), ShouldResemble, Entry{"level": "panic", "message": "Panicf"})
			So(q.Pop(), ShouldResemble, Entry{"level": "panic", "message": "Panicln"})
		})
	})
}

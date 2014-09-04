package logr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDiscard(t *testing.T) {
	Convey("Given a Discard logger", t, func() {
		log := &Discard{}

		Convey("It can be used as a Logger", func() {
			_, ok := interface{}(log).(Logger)
			So(ok, ShouldEqual, true)
		})

		Convey("It can be used as a BatchLogger", func() {
			_, ok := interface{}(log).(BatchLogger)
			So(ok, ShouldEqual, true)
		})
	})
}

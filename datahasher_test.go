package datahasher_test

import (
	"testing"

	"git.llnw.com/dcrosby/datahasher"
	"github.com/cespare/xxhash"
	. "github.com/smartystreets/goconvey/convey"
)

func TestComputeHash(t *testing.T) {
	Convey("for a simple string", t, func() {
		str := "hello world!"
		Convey("returns the xxhash Sum64String result", func() {
			So(datahasher.ComputeHash(str), ShouldEqual, xxhash.Sum64String(str))
		})
	})
}

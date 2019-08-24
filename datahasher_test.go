package datahasher_test

import (
	"testing"

	"github.com/cespare/xxhash"
	. "github.com/smartystreets/goconvey/convey"
	"github.llnw.net/dcrosby/datahasher.git"
)

func TestComputeHash(t *testing.T) {
	Convey("for a simple string", t, func() {
		input := "hello world!"
		Convey("returns the xxhash Sum64String result", func() {
			So(datahasher.ComputeHash(input), ShouldEqual, xxhash.Sum64String(input))
		})
	})

	Convey("for a basic struct", t, func() {
		type basicStruct struct {
			Name string
			Age  int
		}
		input := basicStruct{Name: "George", Age: 37}
		hash := datahasher.ComputeHash(input)
		Convey("returns non zero", func() {
			So(hash, ShouldNotBeZeroValue)
		})
		Convey("returns consistent value (when using same struct)", func() {
			So(hash, ShouldEqual, datahasher.ComputeHash(input))
		})
		Convey("returns consistent value (when using ptr to struct)", func() {
			So(hash, ShouldEqual, datahasher.ComputeHash(&input))
		})
		Convey("returns consistent value (when using copied struct)", func() {
			copied := input
			So(hash, ShouldEqual, datahasher.ComputeHash(copied))
		})
		Convey("returns consistent value (when using identical struct)", func() {
			dupe := basicStruct{Name: "George", Age: 37}
			So(hash, ShouldEqual, datahasher.ComputeHash(dupe))
		})
	})

	Convey("for a messy nested struct", t, func() {
		type typeDesc string
		type idBase struct {
			id   uint64
			Type typeDesc
		}
		type intBox struct {
			Value int
		}
		type messyStruct struct {
			idBase
			Nums  []*intBox
			Lists map[string][]*intBox
		}

		input := messyStruct{
			idBase: idBase{
				id:   42,
				Type: typeDesc("myType"),
			},
			Nums: []*intBox{
				&intBox{Value: 11},
				&intBox{Value: 22},
			},
			Lists: map[string][]*intBox{
				"first": []*intBox{
					&intBox{Value: 33},
					&intBox{Value: 44},
				},
				"second": []*intBox{
					&intBox{Value: 55},
					&intBox{Value: 66},
				},
			},
		}

		hash := datahasher.ComputeHash(input)
		Convey("returns non zero", func() {
			So(hash, ShouldNotBeZeroValue)
		})
		Convey("returns consistent value (when using same struct)", func() {
			So(hash, ShouldEqual, datahasher.ComputeHash(input))
		})
		Convey("returns consistent value (when using ptr to struct)", func() {
			So(hash, ShouldEqual, datahasher.ComputeHash(&input))
		})
		// Convey("returns consistent value (when using copied struct)", func() {
		// 	copied := input
		// 	So(hash, ShouldEqual, datahasher.ComputeHash(copied))
		// })
		Convey("when using equivalent struct", func() {
			dupe := messyStruct{
				idBase: idBase{
					id:   42,
					Type: typeDesc("myType"),
				},
				Nums: []*intBox{
					&intBox{Value: 11},
					&intBox{Value: 22},
				},
				Lists: map[string][]*intBox{
					"first": []*intBox{
						&intBox{Value: 33},
						&intBox{Value: 44},
					},
					"second": []*intBox{
						&intBox{Value: 55},
						&intBox{Value: 66},
					},
				},
			}
			Convey("returns equal value", func() {
				So(hash, ShouldEqual, datahasher.ComputeHash(dupe))
			})
			Convey("...and a nested value gets tweaked", func() {
				dupe.Nums[1].Value = 23
				Convey("returns different hash", func() {
					So(hash, ShouldNotEqual, datahasher.ComputeHash(dupe))
				})
			})
		})
	})

	Convey("for custom DataHasher implementors", t, func() {
		liar1 := Liar{Value: 42}
		liar2 := Liar{Value: 37}
		Convey("returns custom hash value despite struct content", func() {
			So(datahasher.ComputeHash(liar1), ShouldEqual, 424242)
			So(datahasher.ComputeHash(liar2), ShouldEqual, 424242)
		})

	})

}

type Liar struct {
	Value int
}

func (me Liar) DataHash() uint64 {
	return 424242
}

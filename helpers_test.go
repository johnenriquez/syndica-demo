package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is doubled", func() {
			x *= 2

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func TestGetTimeAgo(t *testing.T) {
	cases := []struct{ in, exp string }{
		{"2018-08-10 06:38:57", "8 days ago"},
		{"2018-08-09 06:38:57", "9 days ago"},
	}
	for _, c := range cases {
		out := GetTimeAgo(c.in)
		if out != c.exp {
			t.Errorf("not right, in %s, out %s, exp %s", c.in, out, c.exp)
		}

	}
}

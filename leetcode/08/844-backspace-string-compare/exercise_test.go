package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		s      string
		t      string
		output bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
		{"a#c", "b", false},
	}

	implementations = []struct {
		name string
		f    func(s string, t string) bool
	}{
		{"backspaceCompare", backspaceCompare},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(ca.s, ca.t)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					f(c.s, c.t)
				}
			}
		})
	}
}

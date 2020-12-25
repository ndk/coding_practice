package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		text    string
		pattern string
		output  bool
	}{
		{"a", "a", true},
		{"aa", "a", false},
		{"aa", "aa", true},
		{"abc", "a.c", true},
		{"abbb", "ab*", true},
		{"acd", "ab*c.", true},
	}

	implementations = []struct {
		name string
		f    func(text string, pattern string) bool
	}{
		{"isMatch", isMatch},
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
						result := f(ca.text, ca.pattern)
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
					f(c.text, c.pattern)
				}
			}
		})
	}
}

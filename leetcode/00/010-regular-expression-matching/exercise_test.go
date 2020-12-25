package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		s      string
		p      string
		output bool
	}{
		{"", "", true},
		{"a", "a*", true},
		{"a", "a", true},
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"abc", "a.c", true},
		{"abbb", "ab*", true},
		{"acd", "ab*c.", true},
		{"aab", "c*a*b", true},
		{"b", "c*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"a", "ab*a", false},
		{"a", ".*..a*", false},
	}

	implementations = []struct {
		name string
		f    func(s string, p string) bool
	}{
		{"isMatch", isMatch},
		{"isMatch1", isMatch1},
		{"isMatch2", isMatch2},
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
						result := f(ca.s, ca.p)
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
					f(c.s, c.p)
				}
			}
		})
	}
}

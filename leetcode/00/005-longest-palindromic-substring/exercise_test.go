package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		s      string
		output []string
	}{
		{"", []string{""}},
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"ac", []string{"a", "c"}},
	}

	implementations = []struct {
		name string
		f    func(s string) string
	}{
		{"longestPalindrome", longestPalindrome},
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
						result := f(ca.s)
						So(result, ShouldBeIn, ca.output)
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
					f(c.s)
				}
			}
		})
	}
}

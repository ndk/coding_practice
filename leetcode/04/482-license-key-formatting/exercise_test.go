package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		s      string
		k      int
		output string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	}

	implementations = []struct {
		name string
		f    func(s string, k int) string
	}{
		{"licenseKeyFormatting", licenseKeyFormatting},
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
						result := f(ca.s, ca.k)
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
					f(c.s, c.k)
				}
			}
		})
	}
}

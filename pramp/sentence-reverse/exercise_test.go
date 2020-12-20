package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  string
		output string
	}{
		{"perfect makes practice", "practice makes perfect"},
	}

	implementations = []struct {
		name string
		f    func(sentence string) string
	}{
		{"reverseSentence", reverseSentence},
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
						result := f(ca.input)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func Benchmark_Exercise(b *testing.B) {
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					f(c.input)
				}
			}
		})
	}
}

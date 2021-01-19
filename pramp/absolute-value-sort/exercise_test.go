package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  []int
		output []int
	}{
		{[]int{2, -7, -2, -2, 0}, []int{0, -2, -2, 2, -7}},
	}

	implementations = []struct {
		name string
		f    func(arr []int)
	}{
		{"absSort", absSort},
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
						input := make([]int, len(ca.input))
						copy(input, ca.input)
						f(input)
						So(input, ShouldResemble, ca.output)
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
					b.StopTimer()
					input := make([]int, len(c.input))
					copy(input, c.input)
					b.StartTimer()
					f(input)
				}
			}
		})
	}
}

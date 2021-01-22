package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  []int
		s      int
		output []int
	}{
		{[]int{2, 7, 4, 0, 9, 5, 1, 3}, 20, []int{0, 4, 7, 9}},
	}

	implementations = []struct {
		name string
		f    func(arr []int, s int) []int
	}{
		{"findArrayQuadruplet", findArrayQuadruplet},
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
						f(input, ca.s)
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
					f(input, c.s)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  []int
		output int
	}{
		{[]int{0, 1, 2, 3}, 4},
		{[]int{0, 1, 3}, 2},
		{[]int{1, 2, 3}, 0},
		{[]int{0, 2, 4}, 1},
	}

	implementations = []struct {
		name string
		f    func(input []int) int
	}{
		{"getDifferentNumber", getDifferentNumber},
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

func BenchmarkExercise(b *testing.B) {
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

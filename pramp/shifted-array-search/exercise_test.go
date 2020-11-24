package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		shiftArr []int
		num      int
		output   int
	}{
		{[]int{2}, 2, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{0, 1, 2, 3, 4, 5}, 1, 1},
		{[]int{1, 2, 3, 4, 5, 0}, 0, 5},
		{[]int{9, 12, 17, 2, 4, 5}, 17, 2},
		{[]int{9, 12, 17, 2, 4, 5, 6}, 4, 4},
		{[]int{9, 12, 17, 2, 4, 5}, 2, 3},
		{[]int{9, 12, 17, 2, 4, 5}, 3, -1},
	}

	implementations = []struct {
		name string
		f    func(shiftArr []int, num int) int
	}{
		{"ShiftedArrSearch", ShiftedArrSearch},
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
						result := f(ca.shiftArr, ca.num)
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
					f(c.shiftArr, c.num)
				}
			}
		})
	}
}

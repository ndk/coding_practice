package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		nums   []int
		output int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{-1}, -1},
		{[]int{-2, -1}, -1},
		{[]int{-1, -2}, -1},
		{[]int{-2, 1}, 1},
		{[]int{-2147483647}, -2147483647},
	}

	implementations = []struct {
		name string
		f    func(nums []int) int
	}{
		{"maxSubArray", maxSubArray},
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
						result := f(ca.nums)
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
					f(c.nums)
				}
			}
		})
	}
}

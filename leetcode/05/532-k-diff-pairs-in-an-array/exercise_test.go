package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
		{[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3, 2},
		{[]int{-1, -2, -3}, 1, 2},
		{[]int{1, 1, 1, 1, 1}, 0, 1},
	}

	implementations = []struct {
		name string
		f    func(nums []int, k int) int
	}{
		{"findPairs", findPairs},
		{"findPairs2", findPairs2},
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
						result := f(ca.nums, ca.k)
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
					f(c.nums, c.k)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		nums   []int
		target int
		output [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nil,
			0,
			[][]int(nil),
		},
	}

	implementations = []struct {
		name string
		f    func(nums []int, target int) [][]int
	}{
		{"fourSum", fourSum},
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
						nums := make([]int, len(ca.nums))
						copy(nums, ca.nums)
						result := f(nums, ca.target)
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
					b.StopTimer()
					nums := make([]int, len(c.nums))
					copy(nums, c.nums)
					b.StartTimer()
					f(nums, c.target)
				}
			}
		})
	}
}

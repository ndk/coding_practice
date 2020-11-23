package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		arr    []int
		output int
	}{
		{[]int{-8, 0, 2, 5}, 2},
		{[]int{-1, 0, 3, 6}, -1},
		{[]int{-6, -5, -4, -1, 1, 3, 5, 7}, 7},
	}

	implementations = []struct {
		name string
		f    func(arr []int) int
	}{
		{"IndexEqualsValueSearch", IndexEqualsValueSearch},
		{"IndexEqualsValueSearch_linear", IndexEqualsValueSearch_linear},
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
						result := f(ca.arr)
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
					f(c.arr)
				}
			}
		})
	}
}

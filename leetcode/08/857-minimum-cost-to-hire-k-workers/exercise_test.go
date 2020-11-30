package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		quality []int
		wage    []int
		k       int
		output  float64
	}{
		{[]int{10, 20, 5}, []int{70, 50, 30}, 2, 105},
		{[]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3, 30.66667},
		{[]int{25, 68, 35, 62, 52, 57, 35, 83, 40, 51}, []int{147, 97, 251, 129, 438, 443, 120, 366, 362, 343}, 6, 1979.31429},
		{[]int{4, 4, 4, 5}, []int{13, 12, 13, 12}, 2, 26},
		{[]int{5, 4, 1, 3}, []int{4, 3, 14, 9}, 1, 3},
	}

	implementations = []struct {
		name string
		f    func(quality []int, wage []int, k int) float64
	}{
		{"mincostToHireWorkers", mincostToHireWorkers},
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
						result := f(ca.quality, ca.wage, ca.k)
						So(result, ShouldAlmostEqual, ca.output, .00001)
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
					f(c.quality, c.wage, c.k)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		numCourses    int
		prerequisites [][]int
		output        []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{0, 1}}, []int{1, 0}},
		{2, [][]int{{0, 1}, {1, 0}}, []int(nil)},
	}

	implementations = []struct {
		name string
		f    func(numCourses int, prerequisites [][]int) []int
	}{
		{"findOrder", findOrder},
		{"findOrderBFS", findOrderBFS},
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
						result := f(ca.numCourses, ca.prerequisites)
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
					f(c.numCourses, c.prerequisites)
				}
			}
		})
	}
}

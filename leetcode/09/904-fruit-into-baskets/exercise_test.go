package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		tree   []int
		output int
	}{
		{[]int{1}, 1},
		{[]int{1, 1}, 2},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
	}

	implementations = []struct {
		name string
		f    func(tree []int) int
	}{
		{"totalFruit", totalFruit},
		{"totalFruit2", totalFruit2},
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
						result := f(ca.tree)
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
					f(c.tree)
				}
			}
		})
	}
}

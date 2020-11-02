package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		a      []int
		output int
	}{
		{[]int{10, 13, 12, 14, 15}, 2},
		{[]int{2, 3, 1, 1, 4}, 3},
		{[]int{5, 1, 3, 4, 2}, 3},
		{[]int{1, 2, 3, 2, 1, 4, 4, 5}, 6},
		{[]int{81, 54, 96, 60, 58}, 2},
	}

	implementations = []struct {
		name string
		f    func(a []int) int
	}{
		{"oddEvenJumps", oddEvenJumps},
		{"oddEvenJumps2", oddEvenJumps2},
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
						result := f(ca.a)
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
					f(c.a)
				}
			}
		})
	}
}

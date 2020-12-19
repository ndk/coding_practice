package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		grantsArray []int
		newBudget   int
		output      float32
	}{
		{[]int{}, 10, 0},
		{[]int{2}, 10, 2},
		{[]int{10}, 10, 10},
		{[]int{2, 8}, 10, 8},
		{[]int{2, 7}, 10, 7},
		{[]int{2, 100, 50, 120, 1000}, 190, 47},
		{[]int{100, 300, 200, 400}, 800, 250},
		{[]int{14, 15, 16, 17, 18, 19}, 97, 17.5},
	}

	implementations = []struct {
		name string
		f    func(grantsArray []int, newBudget int) float32
	}{
		{"findGrantsCap", findGrantsCap},
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
						result := f(ca.grantsArray, ca.newBudget)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func Benchmark_Exercise(b *testing.B) {
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					f(c.grantsArray, c.newBudget)
				}
			}
		})
	}
}

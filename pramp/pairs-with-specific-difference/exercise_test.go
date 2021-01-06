package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		arr    []int
		k      int
		output [][2]int
	}{
		{
			[]int{0, -1, -2, 2, 1},
			1,
			[][2]int{{1, 0}, {0, -1}, {-1, -2}, {2, 1}},
		},
		{
			[]int{1, 7, 5, 3, 32, 17, 12},
			17,
			[][2]int(nil),
		},
	}

	implementations = []struct {
		name string
		f    func(arr []int, k int) [][2]int
	}{
		{"findPairsWithGivenDifference", findPairsWithGivenDifference},
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
						result := f(ca.arr, ca.k)
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
					f(c.arr, c.k)
				}
			}
		})
	}
}

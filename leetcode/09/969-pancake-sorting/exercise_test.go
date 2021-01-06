package exercise

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		arr    []int
		output []int
	}{
		{[]int{3, 2, 4, 1}, []int{3, 4, 2, 3, 2}},
		{[]int{1, 2, 3}, []int{3, 3, 2, 2}},
	}

	implementations = []struct {
		name string
		f    func(arr []int) []int
	}{
		{"pancakeSort", pancakeSort},
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
						arr := make([]int, len(ca.arr))
						copy(arr, ca.arr)
						result := f(arr)
						So(sort.IntsAreSorted(arr), ShouldBeTrue)
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

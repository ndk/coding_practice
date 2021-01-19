package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  []int
		output []int
	}{
		{[]int{1, 5, 4, 3, 2}, []int{1, 2, 3, 4, 5}},
	}

	implementations = []struct {
		name string
		f    func(arr []int)
	}{
		{"pancakeSort", pancakeSort},
	}
)

func TestFlip(t *testing.T) {
	Convey("Test", t, func() {
		cases := []struct {
			input  []int
			k      int
			output []int
		}{
			{[]int{}, 0, []int{}},
			{[]int{1}, 0, []int{1}},
			{[]int{1, 2}, 0, []int{1, 2}},
			{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
			{[]int{1}, 1, []int{1}},
			{[]int{1, 2}, 1, []int{1, 2}},
			{[]int{1, 2, 3}, 1, []int{1, 2, 3}},
			{[]int{1, 2}, 2, []int{2, 1}},
			{[]int{1, 2, 3}, 2, []int{2, 1, 3}},
			{[]int{1, 2, 3}, 3, []int{3, 2, 1}},
			{[]int{1, 2, 3, 4}, 3, []int{3, 2, 1, 4}},
		}
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				input := make([]int, len(ca.input))
				copy(input, ca.input)
				flip(input, ca.k)
				So(input, ShouldResemble, ca.output)
			})
		}
	})
}

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						input := make([]int, len(ca.input))
						copy(input, ca.input)
						f(input)
						So(input, ShouldResemble, ca.output)
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
					b.StopTimer()
					input := make([]int, len(c.input))
					copy(input, c.input)
					b.StartTimer()
					f(input)
				}
			}
		})
	}
}

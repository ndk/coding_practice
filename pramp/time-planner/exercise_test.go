package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		slotsA [][]int
		slotsB [][]int
		dur    int
		output []int
	}{
		{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 8, []int{60, 68}},
		{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 12, []int{}},
		{[][]int{{1, 10}}, [][]int{{2, 3}, {5, 7}}, 2, []int{5, 7}},
	}

	implementations = []struct {
		name string
		f    func(slotsA [][]int, slotsB [][]int, dur int) []int
	}{
		{"MeetingPlanner", MeetingPlanner},
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
						result := f(ca.slotsA, ca.slotsB, ca.dur)
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
					f(c.slotsA, c.slotsB, c.dur)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		data   [][3]int
		output int
	}{
		{
			[][3]int{
				{1487799425, 14, 1},
				{1487799425, 4, 0},
				{1487799425, 2, 0},
				{1487800378, 10, 1},
				{1487801478, 18, 0},
				{1487801478, 18, 1},
				{1487901013, 1, 0},
				{1487901211, 7, 1},
				{1487901211, 7, 0},
			},
			1487800378,
		},
	}

	implementations = []struct {
		name string
		f    func(data [][3]int) int
	}{
		{"findBusiestPeriod", findBusiestPeriod},
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
						result := f(ca.data)
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
					f(c.data)
				}
			}
		})
	}
}

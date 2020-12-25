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
				{0, 2, 10},
				{3, 5, 0},
				{9, 20, 6},
				{10, 12, 15},
				{10, 10, 8},
			},
			5,
		},
	}

	implementations = []struct {
		name string
		f    func(data [][3]int) int
	}{
		{"calcDroneMinEnergy", calcDroneMinEnergy},
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

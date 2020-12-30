package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		input  []string
		output int
	}{
		{
			[]string{
				"01010",
				"00111",
				"10010",
				"01100",
				"10101",
			}, 6,
		},
	}

	implementations = []struct {
		name string
		f    func(matrix [][]byte) int
	}{
		{"getNumberOfIslands", getNumberOfIslands},
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
						matrix := make([][]byte, len(ca.input))
						for i, row := range ca.input {
							matrix[i] = []byte(row)
						}
						result := f(matrix)
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
					b.StopTimer()
					matrix := make([][]byte, len(c.input))
					for i, row := range c.input {
						matrix[i] = []byte(row)
					}
					b.StartTimer()
					f(matrix)
				}
			}
		})
	}
}

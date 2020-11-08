package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		output    []float64
	}{
		{
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
			},
			values: []float64{2.0, 3.0},
			queries: [][]string{
				{"a", "c"},
				{"b", "a"},
				{"a", "e"},
				{"a", "a"},
				{"x", "x"},
			},
			output: []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
				{"bc", "cd"},
			},
			values: []float64{1.5, 2.5, 5.0},
			queries: [][]string{
				{"a", "c"},
				{"c", "b"},
				{"bc", "cd"},
				{"cd", "bc"},
			},
			output: []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			equations: [][]string{
				{"a", "b"},
			},
			values: []float64{0.5},
			queries: [][]string{
				{"a", "b"},
				{"b", "a"},
				{"a", "c"},
				{"x", "y"},
			},
			output: []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}

	implementations = []struct {
		name string
		f    func(equations [][]string, values []float64, queries [][]string) []float64
	}{
		{"calcEquation", calcEquation},
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
						result := f(ca.equations, ca.values, ca.queries)
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
					f(c.equations, c.values, c.queries)
				}
			}
		})
	}
}

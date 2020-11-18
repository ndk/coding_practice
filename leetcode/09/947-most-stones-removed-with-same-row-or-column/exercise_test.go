package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		stones [][]int
		output int
	}{
		{
			[][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 2},
				{2, 1},
				{2, 2},
			},
			5,
		},
		{
			[][]int{
				{0, 0},
				{0, 2},
				{1, 1},
				{2, 0},
				{2, 2},
			},
			3,
		},
		{
			[][]int{
				{0, 0},
			},
			0,
		},
		{
			[][]int{
				{0, 0},
				{1, 0},
			},
			1,
		},
		{
			[][]int{
				{0, 0},
				{1, 0},
				{2, 0},
			},
			2,
		},
		{
			[][]int{
				{0, 1},
				{0, 0},
				{1, 0},
				{2, 0},
			},
			3,
		},
		{
			[][]int{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
			},
			3,
		},
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			0,
		},
		{
			[][]int{
				{3, 2},
				{3, 1},
				{4, 4},
				{1, 1},
				{0, 2},
				{4, 0},
			},
			4,
		},
	}

	implementations = []struct {
		name string
		f    func(stones [][]int) int
	}{
		{"removeStones", removeStones},
		{"removeStones_leet", removeStones_leet},
		{"removeStones_leet2", removeStones_leet2},
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
						result := f(ca.stones)
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
					f(c.stones)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		grid   []string
		output int
	}{
		{
			[]string{
				"11110",
				"11010",
				"11000",
				"00000",
			},
			1,
		},
		{
			[]string{
				"11000",
				"11000",
				"00100",
				"00011",
			},
			3,
		},
		{
			[]string{
				"111111111",
				"100000001",
				"101010101",
				"101010101",
				"101010101",
				"100000001",
				"111111111",
			},
			4,
		},
		{
			[]string{
				"10011101100000000000",
				"10011001000101010010",
				"00011110101100001010",
				"00011001000111001001",
				"00000001110000000000",
				"10000101011000000101",
				"00010001010101010101",
				"00010100110101101110",
				"00001001100001000101",
				"00100100000100100010",
				"10010000000100101010",
				"01000101011011101100",
				"11010000100000010001",
				"01001110001111101000",
				"00111000110001010000",
				"10010100001000101011",
				"10100000010001010000",
				"01100011101010111100",
				"01000011001010010011",
				"00000011110100011000",
			},
			58,
		},
	}

	implementations = []struct {
		name string
		f    func(grid [][]byte) int
	}{
		{"numIslands", numIslands},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				var grid [][]byte
				for _, row := range ca.grid {
					gridRow := make([]byte, len(row))
					for i, b := range row {
						gridRow[i] = byte(b)
					}
					grid = append(grid, gridRow)
				}
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(grid)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	var cs [][][]byte
	for _, c := range cases {
		var grid [][]byte
		for _, row := range c.grid {
			gridRow := make([]byte, len(row))
			for i, b := range row {
				gridRow[i] = byte(b)
			}
			grid = append(grid, gridRow)
		}
		cs = append(cs, grid)
	}
	b.ResetTimer()

	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cs {
					f(c)
				}
			}
		})
	}
}

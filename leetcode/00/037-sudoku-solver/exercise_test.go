package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		board  []string
		output []string
	}{
		{
			[]string{
				"53..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			},
			[]string{
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286179",
			},
		},
	}

	implementations = []struct {
		name string
		f    func(board [][]byte)
	}{
		{"solveSudoku", solveSudoku},
	}
)

func s2b(s []string) [][]byte {
	board := make([][]byte, len(s))
	for i, row := range s {
		board[i] = make([]byte, len(row))
		for j, c := range row {
			board[i][j] = byte(c)
		}
	}
	return board
}

func b2s(b [][]byte) []string {
	s := make([]string, len(b))
	for i, row := range b {
		s[i] = string(row)
	}
	return s
}

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						board := s2b(ca.board)
						f(board)
						result := b2s(board)
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
					board := s2b(c.board)
					b.StartTimer()
					f(board)
				}
			}
		})
	}
}

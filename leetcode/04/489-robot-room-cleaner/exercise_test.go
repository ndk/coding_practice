package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		room     [][]int
		row, col int
	}{
		{
			[][]int{
				{1, 1, 1, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1, 0, 1, 1},
				{1, 0, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
			},
			1, 2,
		},
	}

	implementations = []struct {
		name string
		f    func(robot *Robot)
	}{
		{"cleanRoom", cleanRoom},
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
						robot := Robot{
							row: ca.row,
							col: ca.col,
						}
						robot.room = ca.room
						for _, row := range robot.room {
							cells := make([]int, len(row))
							for i, cell := range row {
								if cell != 0 {
									cells[i] = 1
									robot.remains++
								}
							}
							robot.state = append(robot.state, cells)
						}
						f(&robot)
						So(robot.remains, ShouldBeZeroValue)
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
					robot := Robot{
						row: c.row,
						col: c.col,
					}
					robot.room = c.room
					for _, row := range robot.room {
						cells := make([]int, len(row))
						for i, cell := range row {
							if cell != 0 {
								cells[i] = 1
								robot.remains++
							}
						}
						robot.state = append(robot.state, cells)
					}
					b.StartTimer()
					f(&robot)
				}
			}
		})
	}
}

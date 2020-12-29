package exercise

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		x      float32
		y      float32
		length float32
		depth  int
		lines  [][4]float32
	}{
		{
			0, 0, 8, 2,
			[][4]float32{
				{-4.0, -4.0, -4.0, 4.0},
				{4.0, -4.0, 4.0, 4.0},
				{-4.0, 0.0, 4.0, 0.0},
				{-6.82842712474619, -6.82842712474619, -6.82842712474619, -1.1715728752538102},
				{-1.1715728752538102, -6.82842712474619, -1.1715728752538102, -1.1715728752538102},
				{-6.82842712474619, -4.0, -1.1715728752538102, -4.0},
				{-6.82842712474619, 1.1715728752538102, -6.82842712474619, 6.82842712474619},
				{-1.1715728752538102, 1.1715728752538102, -1.1715728752538102, 6.82842712474619},
				{-6.82842712474619, 4.0, -1.1715728752538102, 4.0},
				{1.1715728752538102, -6.82842712474619, 1.1715728752538102, -1.1715728752538102},
				{6.82842712474619, -6.82842712474619, 6.82842712474619, -1.1715728752538102},
				{1.1715728752538102, -4.0, 6.82842712474619, -4.0},
				{1.1715728752538102, 1.1715728752538102, 1.1715728752538102, 6.82842712474619},
				{6.82842712474619, 1.1715728752538102, 6.82842712474619, 6.82842712474619},
				{1.1715728752538102, 4.0, 6.82842712474619, 4.0},
			},
		},
	}

	implementations = []struct {
		name string
		f    func(x, y float32, length float32, depth int, dl drawLine)
	}{
		{"drawHTree", drawHTree},
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
						lines := [][4]float32{}
						f(ca.x, ca.y, ca.length, ca.depth,
							func(x1, y1, x2, y2 float32) {
								lines = append(lines, [4]float32{x1, y1, x2, y2})
							},
						)
						expectation := make([][4]float32, len(ca.lines))
						copy(expectation, ca.lines)
						for _, slice := range [2][][4]float32{expectation, lines} {
							sort.Slice(slice, func(i, j int) bool {
								a, b := slice[i], slice[j]
								for p := 0; p < len(a); p++ {
									if a[p] != b[p] {
										return a[p] < b[p]
									}
								}
								return false
							})
						}
						So(lines, ShouldResemble, expectation)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	drawLine := func(x1, y1, x2, y2 float32) {}
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					f(c.x, c.y, c.length, c.depth, drawLine)
				}
			}
		})
	}
}

package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	cases = []struct {
		n      int
		edges  [][]int
		output bool
	}{
		{5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},
		{4, [][]int{{2, 3}, {1, 2}, {1, 3}}, false},
	}

	implementations = []struct {
		name string
		f    func(n int, edges [][]int) bool
	}{
		{"validTree", validTree},
		{"validTree_GT_IDFS", validTree_GT_IDFS},
		{"validTree_GT_RDFS", validTree_GT_RDFS},
		{"validTree_GT_BFS", validTree_GT_BFS},
		{"validTree_AGT_IDFS", validTree_AGT_IDFS},
		{"validTree_AGT_RDFS", validTree_AGT_RDFS},
		{"validTree_AGT_BFS", validTree_AGT_BFS},
		{"validTree_UF", validTree_UF},
		{"validTree_UF_PC", validTree_UF_PC},
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
						result := f(ca.n, ca.edges)
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
					f(c.n, c.edges)
				}
			}
		})
	}
}

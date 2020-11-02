package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "cp/leetcode/utils"
)

var (
	cases = [][]int{
		nil,
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4, 5, 6},
	}

	implementations = []struct {
		name string
		f    func(root *BTreeNode) int
	}{
		{"countNodes", countNodes},
		{"countNodes_leet2", countNodes_leet2},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				tree := Ints2BTree(ca)
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(tree)
						So(result, ShouldResemble, len(ca))
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	var trees []*BTreeNode
	for _, c := range cases {
		trees = append(trees, Ints2BTree(c))
	}
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tree := range trees {
					f(tree)
				}
			}
		})
	}
}

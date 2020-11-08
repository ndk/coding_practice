package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "cp/leetcode/utils"
)

var (
	cases = []struct {
		tree   []int
		output int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{3, 1, NULL, NULL, 2}, 2},
	}

	implementations = []struct {
		name string
		f    func(root *BTreeNode) int
	}{
		{"diameterOfBinaryTree", diameterOfBinaryTree},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				tree := Ints2BTree(ca.tree)
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(tree)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	var trees []*BTreeNode
	for _, c := range cases {
		trees = append(trees, Ints2BTree(c.tree))
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

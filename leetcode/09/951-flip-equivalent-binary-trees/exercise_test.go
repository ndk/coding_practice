package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "cp/leetcode/utils"
)

var (
	cases = []struct {
		root1  []int
		root2  []int
		output bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, NULL, NULL, NULL, 7, 8}, []int{1, 3, 2, NULL, 6, 4, 5, NULL, NULL, NULL, NULL, 8, 7}, true},
		{[]int{}, []int{}, true},
		{[]int{}, []int{1}, false},
		{[]int{0, NULL, 1}, []int{}, false},
		{[]int{0, NULL, 1}, []int{0, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6, NULL, NULL, NULL, 7, 8}, []int{99, 3, 2, NULL, 6, 4, 5, NULL, NULL, NULL, NULL, 8, 7}, false},
	}

	implementations = []struct {
		name string
		f    func(root1 *BTreeNode, root2 *BTreeNode) bool
	}{
		{"flipEquiv", flipEquiv},
	}
)

func TestExercise(t *testing.T) {
	Convey("Test", t, func() {
		for i, c := range cases {
			ca := c
			Convey(fmt.Sprintf("case %d", i), func() {
				root1 := Ints2BTree(ca.root1)
				root2 := Ints2BTree(ca.root2)
				for _, impl := range implementations {
					f := impl.f
					Convey(impl.name, func() {
						result := f(root1, root2)
						So(result, ShouldResemble, ca.output)
					})
				}
			})
		}
	})
}

func BenchmarkExercise(b *testing.B) {
	var trees [][2]*BTreeNode
	for _, c := range cases {
		trees = append(trees, [2]*BTreeNode{Ints2BTree(c.root1), Ints2BTree(c.root2)})
	}
	for _, implementation := range implementations {
		f := implementation.f
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tree := range trees {
					f(tree[0], tree[1])
				}
			}
		})
	}
}

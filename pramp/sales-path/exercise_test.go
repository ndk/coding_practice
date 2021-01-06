package exercise

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "cp/leetcode/utils"
)

var (
	cases = []struct {
		tree   string
		output int
	}{
		{"(0(5(4))(3(2(1(1)))(0(10)))(6(1)(5)))", 7},
	}

	implementations = []struct {
		name string
		f    func(tree *TreeNode) int
	}{
		{"getCheapestCost", getCheapestCost},
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
						tree := String2Tree(ca.tree)
						result := f(tree)
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
					tree := String2Tree(c.tree)
					b.StartTimer()
					f(tree)
				}
			}
		})
	}
}

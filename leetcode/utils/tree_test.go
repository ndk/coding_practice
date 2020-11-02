package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString2Tree(t *testing.T) {
	Convey("Test", t, func() {
		Convey("An one-item tree", func() {
			tree := String2Tree("(1)")
			So(tree.Val, ShouldEqual, 1)
			So(tree.Children, ShouldBeEmpty)
		})
		Convey("Children", func() {
			tree := String2Tree("(1(2)(3)(4))")
			So(tree.Val, ShouldEqual, 1)
			So(tree.Children, ShouldHaveLength, 3)
			So(tree.Children[0].Val, ShouldEqual, 2)
			So(tree.Children[1].Val, ShouldEqual, 3)
			So(tree.Children[2].Val, ShouldEqual, 4)
			So(tree.Children[0].Children, ShouldBeEmpty)
			So(tree.Children[1].Children, ShouldBeEmpty)
			So(tree.Children[2].Children, ShouldBeEmpty)
		})
		Convey("A complext tree", func() {
			tree := String2Tree("(1(2(3)(4))(5(6)(7)(8)))")
			So(tree.Val, ShouldEqual, 1)
			So(tree.Children, ShouldHaveLength, 2)
			So(tree.Children[0].Val, ShouldEqual, 2)
			So(tree.Children[1].Val, ShouldEqual, 5)
			So(tree.Children[0].Children[0].Val, ShouldEqual, 3)
			So(tree.Children[0].Children[1].Val, ShouldEqual, 4)
			So(tree.Children[1].Children[0].Val, ShouldEqual, 6)
			So(tree.Children[1].Children[1].Val, ShouldEqual, 7)
			So(tree.Children[1].Children[2].Val, ShouldEqual, 8)
			So(tree.Children[0].Children, ShouldHaveLength, 2)
			So(tree.Children[1].Children, ShouldHaveLength, 3)
		})
		Convey("Panics", func() {
			So(func() {
				String2Tree("1(2(3)(4))(5(6)(7)(8)))")
			}, ShouldPanicWith, `a node definition has been expected`)
			So(func() {
				String2Tree("(1")
			}, ShouldPanicWith, `the node definition is corrupted`)
			So(func() {
				String2Tree("(1a)")
			}, ShouldPanic)
			So(func() {
				String2Tree("(1)1")
			}, ShouldPanicWith, "An incompleted tree definition")
		})
	})
}

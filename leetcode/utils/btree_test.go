package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInts2BTree(t *testing.T) {
	Convey("Test", t, func() {
		Convey("a nil list", func() {
			tree := Ints2BTree(nil)
			So(tree, ShouldBeNil)
		})
		Convey("an empty list", func() {
			tree := Ints2BTree([]int{})
			So(tree, ShouldBeNil)
		})
		Convey("an one-element list", func() {
			tree := Ints2BTree([]int{1})
			So(tree.Val, ShouldEqual, 1)
			So(tree.Left, ShouldBeNil)
			So(tree.Right, ShouldBeNil)
		})
		Convey("a two-elements list", func() {
			Convey("left leaned", func() {
				tree := Ints2BTree([]int{1, 2})
				So(tree.Val, ShouldEqual, 1)
				So(tree.Left.Val, ShouldEqual, 2)
				So(tree.Left.Left, ShouldBeNil)
				So(tree.Left.Right, ShouldBeNil)
				So(tree.Right, ShouldBeNil)
			})
			Convey("right leaned", func() {
				tree := Ints2BTree([]int{1, NULL, 2})
				So(tree.Val, ShouldEqual, 1)
				So(tree.Right.Val, ShouldEqual, 2)
				So(tree.Left, ShouldBeNil)
				So(tree.Right.Left, ShouldBeNil)
				So(tree.Right.Right, ShouldBeNil)
			})
		})
		Convey("a three-elements list", func() {
			tree := Ints2BTree([]int{1, 2, 3})
			So(tree.Val, ShouldEqual, 1)
			So(tree.Left.Val, ShouldEqual, 2)
			So(tree.Right.Val, ShouldEqual, 3)
			So(tree.Left.Left, ShouldBeNil)
			So(tree.Left.Right, ShouldBeNil)
			So(tree.Right.Left, ShouldBeNil)
			So(tree.Right.Right, ShouldBeNil)
		})
		Convey("a complex tree", func() {
			Convey("case 1", func() {
				tree := Ints2BTree([]int{1, 2, 3, NULL, 4, NULL, 5})
				So(tree.Val, ShouldEqual, 1)
				So(tree.Left.Val, ShouldEqual, 2)
				So(tree.Right.Val, ShouldEqual, 3)
				So(tree.Left.Right.Val, ShouldEqual, 4)
				So(tree.Right.Right.Val, ShouldEqual, 5)
				So(tree.Left.Left, ShouldBeNil)
				So(tree.Left.Right.Left, ShouldBeNil)
				So(tree.Left.Right.Right, ShouldBeNil)
				So(tree.Right.Left, ShouldBeNil)
				So(tree.Right.Right.Left, ShouldBeNil)
				So(tree.Right.Right.Right, ShouldBeNil)
			})
		})
	})
}

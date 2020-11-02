package utils

import (
	"strconv"
	"strings"
)

func String2Tree(s string) *TreeNode {
	var f func(s string, node *TreeNode) string
	f = func(s string, node *TreeNode) string {
		if s[0] != '(' {
			panic("a node definition has been expected")
		}
		s = s[1:]

		end := strings.IndexAny(s, "()")
		if end == -1 {
			panic("the node definition is corrupted")
		}
		value, err := strconv.Atoi(s[:end])
		if err != nil {
			panic(err)
		}
		node.Val = value
		s = s[end:]

		for s[0] != ')' {
			var child TreeNode
			node.Children = append(node.Children, &child)
			s = f(s, &child)
		}
		return s[1:]
	}

	var root TreeNode
	if rest := f(s, &root); len(rest) != 0 {
		panic("An incompleted tree definition")
	}

	return &root
}

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

package exercise

import (
	. "cp/leetcode/utils"
)

func getCheapestCost(node *TreeNode) int {
	if len(node.Children) == 0 {
		return node.Val
	}
	if len(node.Children) == 1 {
		return node.Val + getCheapestCost(node.Children[0])
	}

	cheapestSubPath := getCheapestCost(node.Children[0])
	for _, child := range node.Children {
		if cost := getCheapestCost(child); cost < cheapestSubPath {
			cheapestSubPath = cost
		}
	}
	return node.Val + cheapestSubPath
}

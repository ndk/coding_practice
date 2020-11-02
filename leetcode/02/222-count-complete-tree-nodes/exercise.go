package exercise

import (
	. "cp/leetcode/utils"
)

func countNodes(root *BTreeNode) int {
	if root == nil {
		return 0
	}
	left, right, depth := root, root, 0
	for left != nil && right != nil {
		left = left.Left
		right = right.Right
		depth++
	}
	if left == nil && right == nil {
		return 1<<depth - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

///////////////////////////////////////////////////////////////////////////////
// The fastest solution on Leetcode at the moment

func countNodes_leet2(root *BTreeNode) int {
	var explore func(root *BTreeNode, nbNodes int) int
	explore = func(root *BTreeNode, nbNodes int) int {
		if root != nil {
			nbNodes++

			nbNodes = explore(root.Left, nbNodes)
			nbNodes = explore(root.Right, nbNodes)
		}

		return nbNodes
	}

	nbNodes := explore(root, 0)
	return nbNodes
}

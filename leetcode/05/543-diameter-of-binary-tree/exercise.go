package exercise

import (
	. "cp/leetcode/utils"
)

func traverse(root *BTreeNode, longestPath *int) int {
	if root == nil {
		return 0
	}

	forkLength := 0
	longestBranch := 0
	if root.Left != nil {
		left := traverse(root.Left, longestPath)
		longestBranch = left + 1
		forkLength += left + 1
	}
	if root.Right != nil {
		right := traverse(root.Right, longestPath)
		if right+1 > longestBranch {
			longestBranch = right + 1
		}
		forkLength += right + 1
	}
	if forkLength > *longestPath {
		*longestPath = forkLength
	}

	return longestBranch
}

func diameterOfBinaryTree(root *BTreeNode) int {
	var longest int
	traverse(root, &longest)
	return longest
}

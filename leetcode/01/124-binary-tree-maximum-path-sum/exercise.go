package exercise

import (
	"math"

	. "cp/leetcode/utils"
)

func maxPathSum(root *BTreeNode) int {
	var travel func(root *BTreeNode, maxSum *int) int
	travel = func(root *BTreeNode, maxSum *int) int {
		if root == nil {
			return 0
		}

		leftTrack := travel(root.Left, maxSum)
		rightTrack := travel(root.Right, maxSum)
		var longestSubtrack int
		if leftTrack > 0 {
			longestSubtrack = leftTrack
		}
		if rightTrack > longestSubtrack {
			longestSubtrack = rightTrack
		}

		sum := root.Val
		if leftTrack > 0 {
			sum += leftTrack
		}
		if rightTrack > 0 {
			sum += rightTrack
		}
		if sum > *maxSum {
			*maxSum = sum
		}

		return root.Val + longestSubtrack
	}

	maxSum := root.Val
	travel(root, &maxSum)
	return maxSum
}

///////////////////////////////////////////////////////////////////////////////
// The best solution on Leetcode at the moment

func maxPathSum_leetcode(root *BTreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var helper func(root *BTreeNode, maxVal *int) int
	helper = func(root *BTreeNode, maxVal *int) int {
		if root == nil {
			return 0
		}

		l := max(helper(root.Left, maxVal), 0)
		r := max(helper(root.Right, maxVal), 0)

		*maxVal = max(*maxVal, l+r+root.Val)

		return max(l, r) + root.Val
	}

	maxVal := math.MinInt64
	v := helper(root, &maxVal)
	return max(maxVal, v)

}

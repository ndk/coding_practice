package exercise

import (
	. "cp/leetcode/utils"
)

func flipEquiv(root1 *BTreeNode, root2 *BTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	for _, variant := range [][4]*BTreeNode{
		{root1.Left, root1.Right, root2.Left, root2.Right},
		{root1.Left, root1.Right, root2.Right, root2.Left},
	} {
		switch {
		case variant[0] == nil && variant[2] == nil && variant[1] == nil && variant[3] == nil:
			return true
		case variant[0] != nil && variant[2] != nil && variant[1] == nil && variant[3] == nil &&
			variant[0].Val == variant[2].Val:
			return flipEquiv(variant[0], variant[2])
		case variant[0] == nil && variant[2] == nil && variant[1] != nil && variant[3] != nil &&
			variant[1].Val == variant[3].Val:
			return flipEquiv(variant[1], variant[3])
		case variant[0] != nil && variant[2] != nil && variant[1] != nil && variant[3] != nil &&
			variant[0].Val == variant[2].Val && variant[1].Val == variant[3].Val:
			return flipEquiv(variant[0], variant[2]) && flipEquiv(variant[1], variant[3])
		}
	}

	return false
}

package utils

var NULL = -1 << 63

func Ints2BTree(ints []int) *BTreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &BTreeNode{
		Val: ints[0],
	}

	queue := make([]*BTreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &BTreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &BTreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

type BTreeNode struct {
	Val   int
	Left  *BTreeNode
	Right *BTreeNode
}

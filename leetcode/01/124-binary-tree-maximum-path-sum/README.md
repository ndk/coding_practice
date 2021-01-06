### 124. Binary Tree Maximum Path Sum
**Hard**

Given a **non-empty** binary tree, find the maximum path sum.

For this problem, a path is defined as any node sequence from some starting node to any node in the tree along the parent-child connections. The path must contain **at least one node** and does not need to go through the root.

##### Example 1:

  1
 / \
2   3

```
Input: root = [1,2,3]
Output: 6
```

##### Example 2:

 -10
 / \
9  20
  /  \
 15   7

```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
```

##### Note:
* The number of nodes in the tree is in the range `[0, 3 * 104]`.
* `-1000 <= Node.val <= 1000`
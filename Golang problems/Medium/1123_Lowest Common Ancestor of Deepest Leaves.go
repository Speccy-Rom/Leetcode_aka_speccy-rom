/*
Given the root of a binary tree, return the lowest common ancestor of its deepest leaves.

Recall that:

The node of a binary tree is a leaf if and only if it has no children
The depth of the root of the tree is 0. if the depth of a node is d, the depth of each of its children is d + 1.
The lowest common ancestor of a set S of nodes, is the node A with the largest depth such that every node in S is in the subtree with root A.


Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest leaf-nodes of the tree.
Note that nodes 6, 0, and 8 are also leaf nodes, but the depth of them is 2, but the depth of nodes 7 and 4 is 3.
Example 2:

Input: root = [1]
Output: [1]
Explanation: The root is the deepest node in the tree, and it's the lca of itself.
Example 3:

Input: root = [0,1,3,null,2]
Output: [2]
Explanation: The deepest leaf node in the tree is 2, the lca of one node is itself.


Constraints:

The number of nodes in the tree will be in the range [1, 1000].
0 <= Node.val <= 1000
The values of the nodes in the tree are unique.
*/

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type lcaPair struct {
	first  *TreeNode
	second int
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) lcaPair
	dfs = func(root *TreeNode) lcaPair {
		if root == nil {
			return lcaPair{nil, 0}
		}
		l, r := dfs(root.Left), dfs(root.Right)
		d1, d2 := l.second, r.second
		if d1 > d2 {
			return lcaPair{l.first, d1 + 1}
		}
		if d1 < d2 {
			return lcaPair{r.first, d2 + 1}
		}
		return lcaPair{root, d1 + 1}
	}
	return dfs(root).first
}

func testLCA() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	fmt.Println(lcaDeepestLeaves(root))
}

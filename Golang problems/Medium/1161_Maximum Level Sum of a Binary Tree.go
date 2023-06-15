/*
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.
Example 2:

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2


Constraints:

The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105
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
func maxLevelSum(root *TreeNode) int {
	s := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, i int) {
		if root == nil {
			return
		}
		if len(s) == i {
			s = append(s, root.Val)
		} else {
			s[i] += root.Val
		}
		dfs(root.Left, i+1)
		dfs(root.Right, i+1)
	}
	dfs(root, 0)
	ans, mx := 0, -0x3f3f3f3f
	for i, v := range s {
		if mx < v {
			mx = v
			ans = i + 1
		}
	}
	return ans
}

func main() {
	// Test case 1
	// Expected Output: 2
	// root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 7}
	// root.Right = &TreeNode{Val: 0}
	// root.Left.Left = &TreeNode{Val: 7}
	// root.Left.Right = &TreeNode{Val: -8}

	// Test case 2
	// Expected Output: 2
	root := &TreeNode{Val: 989}
	root.Right = &TreeNode{Val: 10250}
	root.Right.Left = &TreeNode{Val: 98693}
	root.Right.Right = &TreeNode{Val: -89388}
	root.Right.Right.Right = &TreeNode{Val: -32127}

	fmt.Println(maxLevelSum(root))
}

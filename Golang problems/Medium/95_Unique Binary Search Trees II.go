/*

Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.



Example 1:


Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
Example 2:

Input: n = 1
Output: [[1]]


Constraints:

1 <= n <= 8

*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTrees(1, n)
}

func generateBSTrees(start, end int) []*TreeNode {
	list := []*TreeNode{}
	if start > end {
		list = append(list, nil)
		return list
	}
	var leftSubtrees, rightSubtrees []*TreeNode
	for i := start; i <= end; i++ {
		leftSubtrees = generateBSTrees(start, i-1)
		rightSubtrees = generateBSTrees(i+1, end)
		for _, left := range leftSubtrees {
			for _, right := range rightSubtrees {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				list = append(list, root)
			}
		}
	}
	return list
}

func main() {
	n := 3
	fmt.Println(generateTrees(n))
}

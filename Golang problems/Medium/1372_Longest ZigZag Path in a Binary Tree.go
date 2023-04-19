/*

 You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.



Example 1:


Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
Example 2:


Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
Example 3:

Input: root = [1]
Output: 0


Constraints:

The number of nodes in the tree is in the range [1, 5 * 104].
1 <= Node.val <= 100
*/

package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// T - структура для хранения максимальных значений
type T struct {
	leftMax    int
	rightMax   int
	subtreeMax int
}

// longestZigZag - функция для нахождения максимальной длины зигзагообразного пути в бинарном дереве
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// dfs - вспомогательная функция для обхода дерева в глубину
	var dfs func(root *TreeNode) T
	dfs = func(root *TreeNode) T {
		if root == nil {
			return T{-1, -1, -1}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		leftZigZag := right.rightMax + 1
		rightZigZag := left.leftMax + 1
		subtreeMax := int(math.Max(float64(leftZigZag), math.Max(float64(rightZigZag), math.Max(float64(left.subtreeMax), float64(right.subtreeMax)))))

		return T{leftZigZag, rightZigZag, subtreeMax}
	}

	return dfs(root).subtreeMax
}

func main() {
	// Пример использования
	// Создаем бинарное дерево
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Left.Right = &TreeNode{Val: 6}
	root.Right.Left.Right.Left = &TreeNode{Val: 7}

	// Вызываем функцию longestZigZag и выводим результат
	result := longestZigZag(root)
	fmt.Println(result) // Вывод: 4
}

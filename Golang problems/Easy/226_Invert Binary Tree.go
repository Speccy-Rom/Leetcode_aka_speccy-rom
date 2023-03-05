/*
Given the root of a binary tree, invert the tree, and return its root.



Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Example 2:


Input: root = [2,1,3]
Output: [2,3,1]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

package main

import "fmt"

// Для решения задачи нам нужно инвертировать двоичное дерево, то есть поменять местами левых и правых потомков каждого узла дерева. Мы можем решить эту задачу с помощью обхода дерева в глубину (DFS) и рекурсивного перехода к левому и правому поддеревьям.
//
//Алгоритм:
//
//Если узел равен nil, вернуть nil.
//Рекурсивно вызвать функцию для левого и правого поддерева.
//Поменять местами левый и правый потомки текущего узла.
//Вернуть измененный узел.
//Так как мы посещаем каждый узел ровно один раз, сложность алгоритма будет O(n), где n - количество узлов в дереве.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func main() {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	fmt.Println(invertTree(root))
}

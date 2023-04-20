/*
Given the root of a binary tree, return the maximum width of the given tree.

The maximum width of a tree is the maximum width among all levels.

The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes), where the null nodes between the end-nodes that would be present in a complete binary tree extending down to that level are also counted into the length calculation.

It is guaranteed that the answer will in the range of a 32-bit signed integer.



Example 1:


Input: root = [1,3,2,5,3,null,9]
Output: 4
Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).
Example 2:


Input: root = [1,3,2,5,null,null,9,6,null,7]
Output: 7
Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).
Example 3:


Input: root = [1,3,2,5]
Output: 2
Explanation: The maximum width exists in the second level with length 2 (3,2).


Constraints:

The number of nodes in the tree is in the range [1, 3000].
-100 <= Node.val <= 100
*/

// Процедура решения задачи будет следующей:
//
//Создать пустую очередь и добавить в нее корневой узел дерева, а также его уровень и позицию (индекс) на этом уровне. Уровень и позиция будут использоваться в дальнейшем для подсчета ширины уровней.
//Создать карту для хранения уровней и их ширин. В качестве ключей карты будем использовать уровни, а в качестве значений - структуру данных, содержащую левую и правую границы позиций узлов на соответствующем уровне. Изначально, левая и правая границы будут равны позиции текущего узла.
//Пока очередь не пуста, продолжать обход в ширину следующим образом:
//a. Извлечь узел из очереди и получить его уровень и позицию.
//b. Обновить границы позиций узлов на текущем уровне, сравнивая их с текущей позицией извлеченного узла.
//c. Добавить в очередь левого и правого потомков извлеченного узла (если они существуют) с их соответствующими уровнями и позициями.
//После завершения обхода, пройти по картам уровней и вычислить максимальную ширину уровней, находя разницу между правой и левой границами позиций узлов на каждом уровне. Вернуть максимальную из полученных ширин.

package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeInfo struct {
	node     *TreeNode
	level    int
	position int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []NodeInfo{{root, 0, 0}}
	levelWidth := make(map[int]*NodeInfo)
	maxWidth := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := levelWidth[node.level]; !ok {
			levelWidth[node.level] = &node
		}

		width := node.position - levelWidth[node.level].position + 1
		maxWidth = max(maxWidth, width)

		if node.node.Left != nil {
			queue = append(queue, NodeInfo{node.node.Left, node.level + 1, node.position * 2})
		}
		if node.node.Right != nil {
			queue = append(queue, NodeInfo{node.node.Right, node.level + 1, node.position*2 + 1})
		}
	}

	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}}
	fmt.Println(widthOfBinaryTree(root))
}

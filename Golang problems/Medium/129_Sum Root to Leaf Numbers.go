/*
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.



Example 1:


Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
Example 2:


Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.


Constraints:

The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 9
The depth of the tree will not exceed 10.

Для решения этой задачи мы можем использовать рекурсию для обхода дерева в глубину (DFS).
На каждом узле мы будем передавать текущее значение числа, которое мы получили по пути от корня к данному узлу.
Если мы дошли до листа, то мы добавляем это число к общей сумме.



При рекурсивном обходе дерева, для каждого узла мы будем обновлять текущее значение числа, умножая его на 10 и добавляя значение узла.
Например, если мы имеем число 12 и переходим к узлу со значением 3, то новое текущее значение числа будет 123 (12 * 10 + 3).

Таким образом, мы можем написать функцию для обхода дерева и вычисления суммы чисел:

*/

package main

import "fmt"

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}
	num = num*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return num
	}
	return dfs(node.Left, num) + dfs(node.Right, num)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(sumNumbers(root))
}

/*
Здесь функция sumNumbers вызывает функцию dfs для обхода дерева, начиная с корня и передает начальное значение числа равное 0.
Функция dfs проверяет, является ли узел листом, и если да, то возвращает текущее значение числа.
В противном случае, она вызывает себя для левого и правого поддерева, передавая обновленное значение числа.
Наконец, сумма значений от листьев добавляется и возвращается.

Временная сложность этого алгоритма составляет O(n), где n - это количество узлов в дереве, потому что мы посещаем каждый узел только один раз.
Пространственная сложность также составляет O(n), так как мы используем рекурсию для обхода дерева.

*/

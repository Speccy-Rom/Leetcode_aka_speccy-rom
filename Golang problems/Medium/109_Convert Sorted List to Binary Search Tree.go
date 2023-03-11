/*
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a
height-balanced
 binary search tree.



Example 1:


Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
Example 2:

Input: head = []
Output: []


Constraints:

The number of nodes in head is in the range [0, 2 * 104].
-105 <= Node.val <= 105

*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMiddle(head)
	root := &TreeNode{Val: mid.Val}
	if head == mid {
		return root
	}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func findMiddle(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}

func main() {
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}
	sortedListToBST(head)
}

/*
Для решения этой задачи мы можем использовать подход "разделяй и властвуй", построив бинарное дерево поиска (BST) из связного списка.
Мы начнем с поиска среднего узла списка, который станет корнем нашего BST.
Затем мы рекурсивно построим левое поддерево из узлов до среднего узла и правое поддерево из узлов после среднего узла.

Для нахождения среднего узла списка мы используем алгоритм "быстрый и медленный указатели".
Мы начинаем с двух указателей на начало списка и перемещаем их через каждый узел списка со скоростью 1 и 2 соответственно.
Когда быстрый указатель достигнет конца списка, медленный указатель будет указывать на средний узел.

Для каждого узла списка мы также храним указатель на предыдущий узел, чтобы мы могли отделить текущий узел от
предыдущего узла при построении BST.

Временная сложность: O(nlogn) - мы рекурсивно обходим каждый узел в списке и каждый узел в дереве, создаваемом из списка.
Общее количество узлов в дереве равно количеству узлов в списке, так что время работы будет O(nlogn).

Пространственная сложность: O(logn) - для каждого вызова рекурсивной функции мы используем константное количество
дополнительной памяти, так что общее количество используемой памяти будет O(logn), где n - количество узлов в списке.

*/

/*

You are given the head of a linked list.

Remove every node which has a node with a greater value anywhere to the right side of it.

Return the head of the modified linked list.



Example 1:


Input: head = [5,2,13,3,8]
Output: [13,8]
Explanation: The nodes that should be removed are 5, 2 and 3.
- Node 13 is to the right of node 5.
- Node 13 is to the right of node 2.
- Node 8 is to the right of node 3.
Example 2:

Input: head = [1,1,1,1]
Output: [1,1,1,1]
Explanation: Every node has value 1, so no nodes are removed.


Constraints:

The number of the nodes in the given list is in the range [1, 105].
1 <= Node.val <= 105

*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	stk := []int{}
	for _, v := range nums {
		for len(stk) > 0 && stk[len(stk)-1] < v {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, v)
	}
	dummy := &ListNode{}
	head = dummy
	for _, v := range stk {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return dummy.Next
}

// Test cases

func main() {
	// Test cases
	res := removeNodes(&ListNode{5, &ListNode{2, &ListNode{13, &ListNode{3, &ListNode{8, nil}}}}})
	for res != nil {
		print(res.Val) // 13 8
		res = res.Next
	}
}

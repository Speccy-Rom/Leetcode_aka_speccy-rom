# You are given two non-empty linked lists representing two non-negative integers.
# The digits are stored in reverse order, and each of their nodes contains a single digit.
# Add the two numbers and return the sum as a linked list.
#
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.
#
#
#
# Example 1:
#
#
# Input: l1 = [2,4,3], l2 = [5,6,4]
# Output: [7,0,8]
# Explanation: 342 + 465 = 807.
# Example 2:
#
# Input: l1 = [0], l2 = [0]
# Output: [0]
# Example 3:
#
# Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
# Output: [8,9,9,9,0,0,0,1]
#
#
# Constraints:
#
# The number of nodes in each linked list is in the range [1, 100].
# 0 <= Node.val <= 9
# It is guaranteed that the list represents a number that does not have leading zeros.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# class Solution:
#     def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
#         digit = 1
#         start_node = ListNode()
#         curr_node = start_node
#         extra = 0
#         while l1 and l2:
#             curr_node.val = (l1.val + l2.val + extra) % 10
#             extra = (l1.val + l2.val + extra) // 10
#             l1 = l1.next
#             l2 = l2.next
#             if l1 and l2:
#                 curr_node.next = ListNode()
#                 curr_node = curr_node.next
#
#         if l1:
#             # while l1:
#             curr_node.next = self.addTwoNumbers(l1, ListNode(val=extra))  # ListNode()
#             extra = 0
#         if l2:
#             curr_node.next = self.addTwoNumbers(l2, ListNode(val=extra))  # ListNode()
#             extra = 0
#         if extra:
#             while curr_node.next:
#                 curr_node = curr_node.next
#             curr_node.next = ListNode(val=extra)
#
#         return start_node


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy = ListNode(0)
        curr = dummy
        carry = 0

        while carry or l1 or l2:
            if l1:
                carry += l1.val
                l1 = l1.next
            if l2:
                carry += l2.val
                l2 = l2.next
            curr.next = ListNode(carry % 10)
            carry //= 10
            curr = curr.next

        return dummy.next


if __name__ == "__main__":
    solution = Solution()
    print(solution.addTwoNumbers([2, 4, 3], [5, 6, 4]))


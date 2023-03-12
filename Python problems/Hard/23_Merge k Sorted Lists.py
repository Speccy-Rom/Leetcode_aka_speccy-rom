# You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
#
# Merge all the linked-lists into one sorted linked-list and return it.
#
#
#
# Example 1:
#
# Input: lists = [[1,4,5],[1,3,4],[2,6]]
# Output: [1,1,2,3,4,4,5,6]
# Explanation: The linked-lists are:
# [
#   1->4->5,
#   1->3->4,
#   2->6
# ]
# merging them into one sorted list:
# 1->1->2->3->4->4->5->6
# Example 2:
#
# Input: lists = []
# Output: []
# Example 3:
#
# Input: lists = [[]]
# Output: []
#
#
# Constraints:
#
# k == lists.length
# 0 <= k <= 104
# 0 <= lists[i].length <= 500
# -104 <= lists[i][j] <= 104
# lists[i] is sorted in ascending order.
# The sum of lists[i].length will not exceed 104.
from typing import List, Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # 1. Собираем все значения в один список
        # 2. Сортируем список
        # 3. Создаем новый список и заполняем его значениями из отсортированного списка
        # 4. Возвращаем новый список
        values = []
        for node in lists:
            while node:
                values.append(node.val)
                node = node.next
        values.sort()
        head = ListNode()
        node = head
        for value in values:
            node.next = ListNode(value)
            node = node.next
        return head.next


if __name__ == '__main__':
    solution = Solution()
    print(solution.mergeKLists([[1, 4, 5], [1, 3, 4], [2, 6]]))
    print(solution.mergeKLists([]))
    print(solution.mergeKLists([[]]))

# Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
#
# There is a cycle in a linked list if there is some node in the list that can be reached again by continuously
# following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is
# connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.
#
# Do not modify the linked list.
#
#
#
# Example 1:
#
#
# Input: head = [3,2,0,-4], pos = 1
# Output: tail connects to node index 1
# Explanation: There is a cycle in the linked list, where tail connects to the second node.
# Example 2:
#
#
# Input: head = [1,2], pos = 0
# Output: tail connects to node index 0
# Explanation: There is a cycle in the linked list, where tail connects to the first node.
# Example 3:
#
#
# Input: head = [1], pos = -1
# Output: no cycle
# Explanation: There is no cycle in the linked list.
from typing import Optional

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                slow = head
                while slow != fast:
                    slow = slow.next
                    fast = fast.next
                return slow
        return None


# Для решения данной задачи, мы можем использовать два указателя, которые двигаются по списку с разной скоростью.

# Алгоритм:

# Создать два указателя slow и fast, которые будут двигаться по списку с разной скоростью
# Пока fast и fast.next не равны None:
# a. Переместить slow на один узел вперед
# b. Переместить fast на два узла вперед
# c. Если slow и fast равны, то:
# i. Переместить slow на начало списка
# ii. Пока slow и fast не равны, переместить slow и fast на один узел вперед
# iii. Вернуть slow
# Вернуть None


if __name__ == '__main__':
    s = Solution()
    print(s.detectCycle([3, 2, 0, -4], 1))
    print(s.detectCycle([1, 2], 0))
    print(s.detectCycle([1], -1))

# Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
#
# Return the smallest level x such that the sum of all the values of nodes at level x is maximal.
# Input: root = [1,7,0,7,-8,null,null]
# Output: 2
# Explanation:
# Level 1 sum = 1.
# Level 2 sum = 7 + 0 = 7.
# Level 3 sum = 7 + -8 = -1.
# So we return the level with the maximum sum which is level 2.
# Example 2:
#
# Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
# Output: 2
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [1, 104].
# -105 <= Node.val <= 105
from collections import deque
from math import inf
from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        q = deque([root])
        mx = -inf
        i = 0
        while q:
            i += 1
            s = 0
            for _ in range(len(q)):
                node = q.popleft()
                s += node.val
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            if mx < s:
                mx = s
                ans = i
        return ans


if __name__ == '__main__':
    s = Solution()
    print(s.maxLevelSum(TreeNode(1, TreeNode(7, TreeNode(7), TreeNode(-8)), TreeNode(0))))
    print(s.maxLevelSum(TreeNode(989, None, TreeNode(10250, TreeNode(98693), TreeNode(-89388, None, TreeNode(-32127))))))
    print(s.maxLevelSum(TreeNode(1)))
    print(s.maxLevelSum(TreeNode(1, TreeNode(2), TreeNode(3))))
    print(s.maxLevelSum(TreeNode(1, TreeNode(2, TreeNode(4)), TreeNode(3, TreeNode(5), TreeNode(6)))))




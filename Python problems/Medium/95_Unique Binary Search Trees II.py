# Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.
#
#
#
# Example 1:
#
#
# Input: n = 3
# Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
# Example 2:
#
# Input: n = 1
# Output: [[1]]
#
#
# Constraints:
#
# 1 <= n <= 8
from typing import List, Optional


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
  def generateTrees(self, n: int) -> List[TreeNode]:
    if n == 0:
      return []

    def generateTrees(mini: int, maxi: int) -> List[Optional[int]]:
      if mini > maxi:
        return [None]

      ans = []

      for i in range(mini, maxi + 1):
        for left in generateTrees(mini, i - 1):
          for right in generateTrees(i + 1, maxi):
            ans.append(TreeNode(i))
            ans[-1].left = left
            ans[-1].right = right

      return ans

    return generateTrees(1, n)



if __name__ == '__main__':
    s = Solution()
    print(s.generateTrees(3))
    print(s.generateTrees(1))

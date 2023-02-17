# Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.
#
#
#
# Example 1:
#
#
# Input: root = [4,2,6,1,3]
# Output: 1
# Example 2:
#
#
# Input: root = [1,0,48,null,null,12,49]
# Output: 1
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [2, 100].
# 0 <= Node.val <= 105


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDiffInBST(self, root: TreeNode) -> int:
        def inorder(node):
            if node:
                yield from inorder(node.left)
                yield node.val
                yield from inorder(node.right)

        prev = None
        res = float('inf')
        for curr in inorder(root):
            if prev is not None:
                res = min(res, abs(curr - prev))
            prev = curr
        return res

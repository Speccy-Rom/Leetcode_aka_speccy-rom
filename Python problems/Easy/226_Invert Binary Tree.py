# Given the root of a binary tree, invert the tree, and return its root.
#
#
#
# Example 1:
#
#
# Input: root = [4,2,7,1,3,6,9]
# Output: [4,7,2,9,6,3,1]
# Example 2:
#
#
# Input: root = [2,1,3]
# Output: [2,3,1]
# Example 3:
#
# Input: root = []
# Output: []
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [0, 100].
# -100 <= Node.val <= 100
#
# # Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return None

        left = self.invertTree(root.left)
        right = self.invertTree(root.right)

        root.left, root.right = right, left

        return root


# Для решения задачи мы рекурсивно вызываем функцию invertTree
# для левого и правого поддеревьев корня, получая развернутые версии этих поддеревьев.
# Затем мы меняем местами левое и правое поддерево у текущего корня и возвращаем этот корень.
#
# Если корень равен None, мы возвращаем None, так как в пустом дереве нечего разворачивать.

if __name__ == "__main__":
    s = Solution()

    # Example 1
    root = TreeNode(4)
    root.left = TreeNode(2)
    root.right = TreeNode(7)
    root.left.left = TreeNode(1)
    root.left.right = TreeNode(3)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(9)
    result = s.invertTree(root)

    # Example 2
    root = TreeNode(2)
    root.left = TreeNode(1)
    root.right = TreeNode(3)
    result = s.invertTree(root)

    # Example 3
    root = None
    result = s.invertTree(root)

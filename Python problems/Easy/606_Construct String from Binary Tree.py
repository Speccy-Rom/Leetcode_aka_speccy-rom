# Given the root of a binary tree, construct a string consisting of parenthesis and integers from a binary tree with the preorder traversal way, and return it.
#
# Omit all the empty parenthesis pairs that do not affect the one-to-one mapping relationship between the string and the original binary tree.
#
#
#
# Example 1:
#
#
# Input: root = [1,2,3,4]
# Output: "1(2(4))(3)"
# Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the unnecessary empty parenthesis pairs. And it will be "1(2(4))(3)"
# Example 2:
#
#
# Input: root = [1,2,3,null,4]
# Output: "1(2()(4))(3)"
# Explanation: Almost the same as the first example, except we cannot omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [1, 104].
# -1000 <= Node.val <= 1000
from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def tree2str(self, root: Optional[TreeNode]) -> str:
        """
        Converts a binary tree to a string representation.

        :param root: The root node of the binary tree.
        :return: The string representation of the binary tree.

        Example:
            Given a binary tree with root node:

                1
               / \
              2   3
             /
            4

            The string representation of the binary tree is "1(2(4))(3)".
        """
        def dfs(root):
            if root is None:
                return ''
            if root.left is None and root.right is None:
                return str(root.val)
            if root.right is None:
                return f'{root.val}({dfs(root.left)})'
            return f'{root.val}({dfs(root.left)})({dfs(root.right)})'

        return dfs(root)


if __name__ == '__main__':
    s = Solution()
    print(s.tree2str(TreeNode(1, TreeNode(2, TreeNode(4)), TreeNode(3))))
    print(s.tree2str(TreeNode(1, TreeNode(2, None, TreeNode(4)), TreeNode(3))))
    print(s.tree2str(TreeNode(1, TreeNode(2, TreeNode(4)), TreeNode(3, TreeNode(5)))))

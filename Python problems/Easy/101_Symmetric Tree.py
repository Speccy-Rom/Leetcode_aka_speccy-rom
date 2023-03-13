#Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
#
#
#
# Example 1:
#
#
# Input: root = [1,2,2,3,4,4,3]
# Output: true
# Example 2:
#
#
# Input: root = [1,2,2,null,3,null,3]
# Output: false
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [1, 1000].
# -100 <= Node.val <= 100
#
#
# Follow up: Could you solve it both recursively and iteratively?


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        def isMirror(t1, t2):
            if not t1 and not t2:
                return True
            if not t1 or not t2:
                return False
            return (t1.val == t2.val) and isMirror(t1.right, t2.left) and isMirror(t1.left, t2.right)
        return isMirror(root.left, root.right)


if __name__ == '__main__':
    s = Solution()
    print(s.isSymmetric([1,2,2,3,4,4,3]))
    print(s.isSymmetric([1,2,2,null,3,null,3]))

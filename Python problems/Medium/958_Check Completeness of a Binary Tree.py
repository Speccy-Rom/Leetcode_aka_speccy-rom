# Given the root of a binary tree, determine if it is a complete binary tree.
#
# In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
#
#
#
# Example 1:
#
#
# Input: root = [1,2,3,4,5,6]
# Output: true
# Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
# Example 2:
#
#
# Input: root = [1,2,3,4,5,null,7]
# Output: false
# Explanation: The node with value 7 isn't as far left as possible.
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [1, 100].
# 1 <= Node.val <= 1000

# Для решения данной задачи мы можем использовать алгоритм обхода дерева в ширину (BFS), который позволит нам посетить все узлы дерева слева направо по уровням.
#
# Для проверки того, что дерево является полным, нам нужно удостовериться, что:
#
# На последнем уровне дерева нет пустых мест между узлами.
# Если какие-то узлы отсутствуют на последнем уровне, они должны быть расположены в самом левом поддереве.
# Мы можем реализовать эту логику следующим образом:
#
# Начнем обход дерева в ширину, используя очередь.
# Для каждого узла проверим, что:
# Если мы уже встретили узел без левого или правого поддерева, все последующие узлы должны быть листьями.
# Если мы уже встретили листовой узел, то все последующие узлы должны быть листьями.
# Если какой-то узел на последнем уровне отсутствует, проверим, что все последующие узлы на том же уровне также отсутствуют.
# Если мы успешно прошли через все узлы, дерево является полным.
# Сложность времени данного решения O(n), где n - количество узлов в дереве.
#
#
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isCompleteTree(self, root: TreeNode) -> bool:
        queue = [root]
        found_empty = False

        while queue:
            node = queue.pop(0)

            # если встретили пустой узел
            if not node:
                found_empty = True

            # если уже встретили узел без левого или правого поддерева
            if found_empty and node:
                return False

            # добавляем узлы в очередь
            if node:
                queue.extend((node.left, node.right))
        return True


# Здесь мы используем очередь для обхода узлов дерева. Переменная found_empty используется для отслеживания того,
# был ли уже найден пустой узел. Если встречается узел без левого или правого поддерева после того, как был найден
# пустой узел, дерево не является полным.
if __name__ == '__main__':
    s = Solution()
    print(s.isCompleteTree([1, 2, 3, 4, 5, 6]))
    print(s.isCompleteTree([1, 2, 3, 4, 5, None, 7]))

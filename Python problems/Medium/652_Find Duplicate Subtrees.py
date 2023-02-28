# Given the root of a binary tree, return all duplicate subtrees.
#
# For each kind of duplicate subtrees, you only need to return the root node of any one of them.
#
# Two trees are duplicate if they have the same structure with the same node values.
#
#
#
# Example 1:
#
#
# Input: root = [1,2,3,4,null,2,4,null,null,4]
# Output: [[2,4],[4]]
# Example 2:
#
#
# Input: root = [2,1,1]
# Output: [[1]]
# Example 3:
#
#
# Input: root = [2,2,2,3,null,3,null]
# Output: [[2,3],[3]]
#
#
# Constraints:
#
# The number of the nodes in the tree will be in the range [1, 5000]
# -200 <= Node.val <= 200
from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        res = {}

        def traverse(node):
            if not node:
                return ''
            left = traverse(node.left)
            right = traverse(node.right)
            sub = f'{str(node.val)},{left},{right}'
            if sub in res and res[sub] == 1:
                res[f'{sub} {id(node)}'] = node
            res[sub] = res.get(sub, 0) + 1
            return sub

        traverse(root)
        return list(res.values())

# Для решения данной задачи, мы можем использовать поиск в глубину (DFS) для обхода всех поддеревьев и сохранения
# каждого поддерева в виде строки, затем используем словарь для хранения всех встретившихся поддеревьев и их количества.
#
# Алгоритм:
#
# Создать пустой словарь res для хранения найденных дубликатов
# Создать функцию traverse, которая принимает на вход текущий узел node:
# a. Если node равен None, вернуть пустую строку
# b. Используя рекурсию, получить значения левого поддерева left и правого поддерева right
# c. Создать строку sub, объединяя значения узла node, left и right через запятую
# d. Если sub уже есть в словаре res и res[sub] равен 1, добавить node в ответ
# e. Добавить sub в словарь res, увеличивая значение на 1
# f. Вернуть sub
# Использовать функцию traverse с корневым узлом root
# Вернуть список всех найденных дубликатов
# Сложность времени: O(n^2) в худшем случае, где n - количество узлов в дереве.
# Сложность пространства: O(n^2) в худшем случае, где n - количество узлов в дереве.


if __name__ == '__main__':
    def test(input1, input2, expected):
        calc = Solution().findDuplicateSubtrees(input1)
        if calc == expected:
            print(f'Test {input1} ok')
        else:
            print(f'Test {input1} fail: {calc} != {expected}')


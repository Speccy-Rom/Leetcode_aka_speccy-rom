# Given the root of a binary tree, return the zigzag level order traversal
# of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
#
#
#
# Example 1:
#
#
# Input: root = [3,9,20,null,null,15,7]
# Output: [[3],[20,9],[15,7]]
# Example 2:
#
# Input: root = [1]
# Output: [[1]]
# Example 3:
#
# Input: root = []
# Output: []
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [0, 2000].
# -100 <= Node.val <= 100


from typing import Optional, List
from collections import deque


# Для решения данной задачи можно использовать алгоритм обхода в ширину (BFS) с некоторыми модификациями.
# В частности, мы будем использовать очередь для хранения узлов, которые нужно посетить на
# текущем уровне, и флаг для обозначения направления обхода на текущем уровне
# (слева направо или справа налево).

# Алгоритм будет проходить все уровни дерева, начиная с корня, добавляя узлы каждого уровня в очередь.
# На каждом уровне мы будем менять направление обхода,
# чтобы получить зигзагообразный порядок значений узлов.
# Как только очередь опустеет, мы закончим обход.

# Таким образом, мы получим двумерный список значений узлов дерева, расположенных в зигзагообразном порядке.
#
# Сложность данного алгоритма по времени составляет O(n), где n - это количество узлов в дереве,
# так как мы посещаем каждый узел ровно один раз.
# По памяти мы используем O(w), где w - это максимальная ширина дерева,
# то есть количество узлов на наибольшем уровне дерева.


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []

        result = []
        queue = deque([root])
        reverse = False

        while queue:
            level = []
            level_size = len(queue)

            for _ in range(level_size):
                node = queue.popleft()
                level.append(node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            if reverse:
                level.reverse()

            result.append(level)
            reverse = not reverse

        return result

# Здесь мы сначала проверяем, не является ли корень дерева пустым. Если это так, мы возвращаем пустой список.
#
# Затем мы создаем очередь, добавляем корень в нее и устанавливаем флаг направления обхода на False (слева направо).
#
# Затем мы начинаем цикл по очереди, извлекая узлы из очереди и добавляя их значения в текущий уровень.
# Если у узла есть левый или правый потомок, мы добавляем их в очередь для посещения позже.
#
# Как только мы посетили все узлы на текущем уровне, мы проверяем флаг направления обхода и,
# если он равен True (справа налево), мы переворачиваем текущий уровень. Затем мы добавляем этот

# You are given the root of a binary tree containing digits from 0 to 9 only.
#
# Each root-to-leaf path in the tree represents a number.
#
# For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
# Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
#
# A leaf node is a node with no children.
#
#
#
# Example 1:
#
#
# Input: root = [1,2,3]
# Output: 25
# Explanation:
# The root-to-leaf path 1->2 represents the number 12.
# The root-to-leaf path 1->3 represents the number 13.
# Therefore, sum = 12 + 13 = 25.
# Example 2:
#
#
# Input: root = [4,9,0,5,1]
# Output: 1026
# Explanation:
# The root-to-leaf path 4->9->5 represents the number 495.
# The root-to-leaf path 4->9->1 represents the number 491.
# The root-to-leaf path 4->0 represents the number 40.
# Therefore, sum = 495 + 491 + 40 = 1026.
#
#
# Constraints:
#
# The number of nodes in the tree is in the range [1, 1000].
# 0 <= Node.val <= 9
# The depth of the tree will not exceed 10.
#
# Для решения данной задачи, мы можем использовать рекурсивный подход для прохода через каждый узел дерева,
# а также для вычисления суммы всех путей от корня до листьев.
#
# Для каждого узла, мы будем проходить через его левое и правое поддеревья, передавая текущее число, образованное из
# предыдущих узлов и текущего узла. Мы будем складывать результаты, полученные от левого и правого поддеревьев,
# и возвращать их в качестве итоговой суммы для данного узла.
#
# Мы будем использовать префиксный метод для вычисления текущего числа, то есть когда мы переходим от одного узла к
# другому, мы умножаем текущее число на 10 и добавляем значение следующего узла.
#
# Для вычисления сложности алгоритма, нам нужно учитывать количество узлов в дереве, которое может быть до 1000,
# и глубину дерева, которая не превышает 10. Таким образом, сложность алгоритма будет O(1000 * 2^10),
# что эквивалентно O(1024000), что является довольно эффективным для данного размера входных данных.

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        return self.dfs(root, 0)

    def dfs(self, node, current_sum):
        if not node:
            return 0

        current_sum = current_sum * 10 + node.val

        if node.left or node.right:
            return self.dfs(node.left, current_sum) + self.dfs(node.right, current_sum)
        else:
            return current_sum


# Мы создали класс TreeNode для представления узла дерева, а также класс Solution для решения задачи. Функция
# sumNumbers принимает корневой узел дерева в качестве аргумента и вызывает функцию dfs для вычисления суммы всех
# путей от корня до листьев.
#
# Функция dfs принимает текущий узел и текущую сумму в качестве аргументов. Мы вычисляем текущую сумму, умножая ее на
# 10 и добавляя значение текущего узла. Затем мы проверяем, является ли текущий узел листом. Если это так,
# мы возвращаем текущую сумму. В противном случае мы вызываем функцию dfs для левого и правого поддеревьев и
# возвращаем сумму результатов.

if __name__ == '__main__':
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    print(Solution().sumNumbers(root))

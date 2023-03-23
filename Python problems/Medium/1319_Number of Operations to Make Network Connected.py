# There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. Any computer can reach any other computer directly or indirectly through the network.
#
# You are given an initial computer network connections. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected.
#
# Return the minimum number of times you need to do this in order to make all the computers connected. If it is not possible, return -1.
#
#
#
# Example 1:
#
#
# Input: n = 4, connections = [[0,1],[0,2],[1,2]]
# Output: 1
# Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
# Example 2:
#
#
# Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
# Output: 2
# Example 3:
#
# Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
# Output: -1
# Explanation: There are not enough cables.
#
#
# Constraints:
#
# 1 <= n <= 105
# 1 <= connections.length <= min(n * (n - 1) / 2, 105)
# connections[i].length == 2
# 0 <= ai, bi < n
# ai != bi
# There are no repeated connections.
# No two computers are connected by more than one cable.


from collections import deque


class Solution:
    def makeConnected(self, n: int, connections: List[List[int]]) -> int:
        if len(connections) < n - 1:
            return -1

        # create adjacency list
        adj_list = [[] for _ in range(n)]
        for u, v in connections:
            adj_list[u].append(v)
            adj_list[v].append(u)

        visited = set()
        components = 0

        for i in range(n):
            if i not in visited:
                components += 1
                visited.add(i)
                queue = deque([i])

                while queue:
                    node = queue.popleft()
                    for neighbor in adj_list[node]:
                        if neighbor not in visited:
                            visited.add(neighbor)
                            queue.append(neighbor)

        return components - 1

# Для решения этой задачи можно использовать алгоритм обхода графа в глубину (DFS) или алгоритм обхода графа в ширину
# (BFS) для нахождения связных компонент графа.
#
# Мы начнем с одной вершины и будем искать все вершины, связанные с ней. Затем мы перейдем к следующей вершине,
# которая еще не была посещена, и найдем все вершины, связанные с ней, и так далее, пока мы не исследуем все вершины
# графа.
#
# Если граф уже полностью связан, то мы можем вернуть 0. В противном случае мы должны добавить ребро между двумя
# компонентами графа, которые не соединены, и продолжать до тех пор, пока все вершины не будут связаны.
#
# Мы будем использовать алгоритм обхода графа в ширину (BFS) для этого. Мы начнем с одной вершины и будем искать все
# вершины, связанные с ней. Затем мы добавим все несвязанные с этой вершиной ребра в очередь, которую мы будем
# использовать для хранения несвязанных компонентов графа.
#
# Когда мы исследуем все вершины, связанные с текущей, мы извлечем первый элемент из очереди и продолжим процесс,
# пока все вершины не будут связаны.
#
# Для определения, связаны ли две вершины графа, мы можем использовать алгоритм обхода графа в глубину (DFS) или
# алгоритм обхода графа в ширину (BFS).
#
# Асимптотическая сложность времени этого алгоритма составляет O(n log n), где n - количество вершин в графе.


if __name__ == '__main__':
    solution = Solution()
    print(solution.makeConnected(4, [[0, 1], [0, 2], [1, 2]]))
    print(solution.makeConnected(6, [[0, 1], [0, 2], [0, 3], [1, 2], [1, 3]]))
    print(solution.makeConnected(6, [[0, 1], [0, 2], [0, 3], [1, 2]]))

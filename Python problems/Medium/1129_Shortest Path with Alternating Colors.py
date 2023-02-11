# You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1. Each edge is red or blue in this graph, and there could be self-edges and parallel edges.
#
# You are given two arrays redEdges and blueEdges where:
#
# redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi in the graph, and
# blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node vj in the graph.
# Return an array answer of length n, where each answer[x] is the length of the shortest path from node 0 to node x such that the edge colors alternate along the path, or -1 if such a path does not exist.
#
#
#
# Example 1:
#
# Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
# Output: [0,1,-1]
# Example 2:
#
# Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
# Output: [0,1,-1]
#
#
# Constraints:
#
# 1 <= n <= 100
# 0 <= redEdges.length, blueEdges.length <= 400
# redEdges[i].length == blueEdges[j].length == 2
# 0 <= ai, bi, uj, vj < n
from collections import deque, defaultdict
from typing import List


# class Solution:
#     def shortestAlternatingPaths(self, n: int, redEdges: List[List[int]], blueEdges: List[List[int]]) -> List[int]:
#         g = defaultdict(list)
#
#         for a, b in redEdges:
#             g[a].append((b, 'red'))
#         for u, v in blueEdges:
#             g[u].append((v, 'blue'))
#
#         answer = [-1 for _ in range(n)]
#
#         q = deque([(0, 'red', 0), (0, 'blue', 0)])  # can start from either blue or red. represents node, color, dist
#         visited = {(0, 'red'), (0, 'blue')}
#
#         answer[0] = 0
#
#         while q:
#             node, color, dist = q.popleft()
#             for new, newcolor in g[node]:
#                 if (new, newcolor) in visited or color == newcolor:
#                     continue
#                 if answer[new] < 0:
#                     answer[new] = dist + 1
#                 q.append((new, newcolor, dist + 1))
#                 visited.add((new, newcolor))
#
#         return answer

class Solution:
    def shortestAlternatingPaths(self, n: int, red_edges: List[List[int]], blue_edges: List[List[int]]) -> List[int]:
        red = [[] for _ in range(n)]
        blue = [[] for _ in range(n)]
        for i, j in red_edges:
            red[i].append(j)
        for i, j in blue_edges:
            blue[i].append(j)
        queue = deque([(0, 0, 0), (0, 1, 0)])
        visited = {(0, 0), (0, 1)}
        res = [-1] * n
        res[0] = 0
        while queue:
            node, color, length = queue.popleft()
            if color == 0:
                for i in red[node]:
                    if (i, 1) not in visited:
                        queue.append((i, 1, length + 1))
                        visited.add((i, 1))
                        res[i] = length + 1 if res[i] == -1 else min(res[i], length + 1)
            else:
                for i in blue[node]:
                    if (i, 0) not in visited:
                        queue.append((i, 0, length + 1))
                        visited.add((i, 0))
                        res[i] = length + 1 if res[i] == -1 else min(res[i], length + 1)
        return res


if __name__ == '__main__':
    n = 3
    red_edges = [[0, 1], [1, 2]]
    blue_edges = []
    print(Solution().shortestAlternatingPaths(n, red_edges, blue_edges))

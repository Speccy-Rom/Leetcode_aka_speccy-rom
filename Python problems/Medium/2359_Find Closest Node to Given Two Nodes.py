# You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.
#
# The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge from node i to node edges[i]. If there is no outgoing edge from i, then edges[i] == -1.
#
# You are also given two integers node1 and node2.
#
# Return the index of the node that can be reached from both node1 and node2, such that the maximum between the distance from node1 to that node, and from node2 to that node is minimized. If there are multiple answers, return the node with the smallest index, and if no possible answer exists, return -1.
#
# Note that edges may contain cycles.
#
#
#
# Example 1:
#
#
# Input: edges = [2,2,3,-1], node1 = 0, node2 = 1
# Output: 2
# Explanation: The distance from node 0 to node 2 is 1, and the distance from node 1 to node 2 is 1.
# The maximum of those two distances is 1. It can be proven that we cannot get a node with a smaller maximum distance than 1, so we return node 2.
# Example 2:
#
#
# Input: edges = [1,2,-1], node1 = 0, node2 = 2
# Output: 2
# Explanation: The distance from node 0 to node 2 is 2, and the distance from node 2 to itself is 0.
# The maximum of those two distances is 2. It can be proven that we cannot get a node with a smaller maximum distance than 2, so we return node 2.
#
#
# Constraints:
#
# n == edges.length
# 2 <= n <= 105
# -1 <= edges[i] < n
# edges[i] != i
# 0 <= node1, node2 < n
from typing import List


# class Solution:
#     def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
#         if node1 == node2:
#             return node1
#
#         graph = defaultdict(set)
#         for i, v in enumerate(edges):
#             if v != -1:
#                 graph[i].add(v)
#
#         def bfs(v):
#             q = deque([(v, 0)])
#             visited = {}
#
#             while q:
#                 node, dist = q.popleft()
#                 visited[node] = dist
#                 for nb in graph[node]:
#                     if nb not in visited:
#                         q.append((nb, dist + 1))
#             return visited
#
#         s1, s2 = bfs(node1), bfs(node2)
#         s = set(s1.keys()) & set(s2.keys())
#         if not s:
#             return -1
#
#         res = -1
#         d = float('inf')
#         for node in s:
#             tmp = max(s1[node], s2[node])
#             if tmp < d:
#                 res = node
#                 d = tmp
#             elif tmp == d:
#                 if node < res:
#                     res = node
#         return res

class Solution:
    def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
        # Increment 1 step at a time, see if there is a matched node using 2 sets.  When there is a match, return it.
        if node1 == node2: return node1

        n1 = {node1}
        n2 = {node2}
        ans = 1000000

        while node1 >= 0 or node2 >= 0:
            e1, e2 = edges[node1], edges[node2]
            if e1 in n1 and e2 in n2: break

            n1.add(e1)
            n2.add(e2)

            if e1 in n2:
                ans = min(ans, e1)
            if e2 in n1:
                ans = min(ans, e2)
            if ans < 1000000:
                return ans

            if e1 != -1: node1 = e1
            if e2 != -1: node2 = e2
        return -1


if __name__ == '__main__':
    print(Solution().closestMeetingNode([2, 2, 3, -1], 0, 1))
    print(Solution().closestMeetingNode([1, 2, -1], 0, 2))

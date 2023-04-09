# There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.
#
# You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.
#
# A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge from xi to xi+1 for every 1 <= i < k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.
#
# Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.
#
#
#
# Example 1:
#
#
#
# Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
# Output: 3
# Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
# Example 2:
#
#
#
# Input: colors = "a", edges = [[0,0]]
# Output: -1
# Explanation: There is a cycle from 0 to 0.
#
#
# Constraints:
#
# n == colors.length
# m == edges.length
# 1 <= n <= 105
# 0 <= m <= 105
# colors consists of lowercase English letters.
# 0 <= aj, bj < n
import collections
from typing import List


class Solution:
    def largestPathValue(self, colors: str, edges: List[List[int]]) -> int:
        n = len(colors)
        ans = 0
        processed = 0
        graph = [[] for _ in range(n)]
        inDegree = [0] * n
        q = collections.deque()
        count = [[0] * 26 for _ in range(n)]

        # Build graph
        for u, v in edges:
            graph[u].append(v)
            inDegree[v] += 1

        # Vpology
        for i, degree in enumerate(inDegree):
            if degree == 0:
                q.append(i)

        while q:
            u = q.popleft()
            processed += 1
            count[u][ord(colors[u]) - ord('a')] += 1
            ans = max(ans, count[u][ord(colors[u]) - ord('a')])
            for v in graph[u]:
                for i in range(26):
                    count[v][i] = max(count[v][i], count[u][i])
                inDegree[v] -= 1
                if inDegree[v] == 0:
                    q.append(v)

        return ans if processed == n else -1


if __name__ == '__main__':
    s = Solution()
    print(s.largestPathValue("abaca", [[0, 1], [0, 2], [2, 3], [3, 4]]))
    print(s.largestPathValue("a", [[0, 0]]))

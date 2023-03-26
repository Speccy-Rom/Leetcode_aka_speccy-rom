# You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.
#
# The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge from node i to node edges[i]. If there is no outgoing edge from node i, then edges[i] == -1.
#
# Return the length of the longest cycle in the graph. If no cycle exists, return -1.
#
# A cycle is a path that starts and ends at the same node.
#
#
#
# Example 1:
#
#
# Input: edges = [3,3,4,2,3]
# Output: 3
# Explanation: The longest cycle in the graph is the cycle: 2 -> 4 -> 3 -> 2.
# The length of this cycle is 3, so 3 is returned.
# Example 2:
#
#
# Input: edges = [2,-1,3,1]
# Output: -1
# Explanation: There are no cycles in this graph.
#
#
# Constraints:
#
# n == edges.length
# 2 <= n <= 105
# -1 <= edges[i] < n
# edges[i] != i
from typing import List


class Solution:
    def longestCycle(self, edges: List[int]) -> int:
        ans = -1
        time = 1
        timeVisited = [0] * len(edges)

        for i, edge in enumerate(edges):
            if timeVisited[i]:
                continue
            startTime = time
            u = i
            while u != -1 and not timeVisited[u]:
                timeVisited[u] = time
                time += 1
                u = edges[u]  # Move to next node
            if u != -1 and timeVisited[u] >= startTime:
                ans = max(ans, time - timeVisited[u])

        return ans


if __name__ == '__main__':
    s = Solution()
    print(s.longestCycle([3, 3, 4, 2, 3]))
    print(s.longestCycle([2, -1, 3, 1]))

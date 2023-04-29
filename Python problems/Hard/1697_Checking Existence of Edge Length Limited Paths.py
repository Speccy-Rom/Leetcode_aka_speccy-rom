# An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi] denotes an edge between nodes ui and vi with distance disi. Note that there may be multiple edges between two nodes.
#
# Given an array queries, where queries[j] = [pj, qj, limitj], your task is to determine for each queries[j] whether there is a path between pj and qj such that each edge on the path has a distance strictly less than limitj .
#
# Return a boolean array answer, where answer.length == queries.length and the jth value of answer is true if there is a path for queries[j] is true, and false otherwise.
#
#
#
# Example 1:
#
#
# Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
# Output: [false,true]
# Explanation: The above figure shows the given graph. Note that there are two overlapping edges between 0 and 1 with distances 2 and 16.
# For the first query, between 0 and 1 there is no path where each distance is less than 2, thus we return false for this query.
# For the second query, there is a path (0 -> 1 -> 2) of two edges with distances less than 5, thus we return true for this query.
# Example 2:
#
#
# Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
# Output: [true,false]
# Exaplanation: The above figure shows the given graph.
#
#
# Constraints:
#
# 2 <= n <= 105
# 1 <= edgeList.length, queries.length <= 105
# edgeList[i].length == 3
# queries[j].length == 3
# 0 <= ui, vi, pj, qj <= n - 1
# ui != vi
# pj != qj
# 1 <= disi, limitj <= 109
# There may be multiple edges between two nodes.


class UnionFind:
    def __init__(self, n: int):
        self.id = list(range(n))
        self.rank = [0] * n

    def union_by_rank(self, u: int, v: int):
        i = self.find(u)
        j = self.find(v)
        if i == j:
            return
        if self.rank[i] < self.rank[j]:
            self.id[i] = j
        elif self.rank[i] > self.rank[j]:
            self.id[j] = i
        else:
            self.id[i] = j
            self.rank[j] += 1

    def find(self, u: int) -> int:
        if self.id[u] != u:
            self.id[u] = self.find(self.id[u])
        return self.id[u]


class Solution:
    def distanceLimitedPathsExist(self, n: int, edgeList: List[List[int]], queries: List[List[int]]) -> List[bool]:
        ans = [False] * len(queries)
        uf = UnionFind(n)

        for i in range(len(queries)):
            queries[i].append(i)

        queries.sort(key=lambda query: query[2])
        edgeList.sort(key=lambda edge: edge[2])

        i = 0  # i := edgeList's index
        for query in queries:
            # Union edges whose distances < limit (query[2])
            while i < len(edgeList) and edgeList[i][2] < query[2]:
                uf.union_by_rank(edgeList[i][0], edgeList[i][1])
                i += 1
            if uf.find(query[0]) == uf.find(query[1]):
                ans[query[3]] = True

        return ans


if __name__ == "__main__":
    import os

    import pytest

    pytest.main([os.path.join("tests", "test_1697.py")])

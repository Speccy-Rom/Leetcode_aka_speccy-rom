# There is an integer array nums that consists of n unique elements, but you have forgotten it. However, you do remember every pair of adjacent elements in nums.
#
# You are given a 2D integer array adjacentPairs of size n - 1 where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.
#
# It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.
#
# Return the original array nums. If there are multiple solutions, return any of them.
#
#
#
# Example 1:
#
# Input: adjacentPairs = [[2,1],[3,4],[3,2]]
# Output: [1,2,3,4]
# Explanation: This array has all its adjacent pairs in adjacentPairs.
# Notice that adjacentPairs[i] may not be in left-to-right order.
# Example 2:
#
# Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
# Output: [-2,4,1,-3]
# Explanation: There can be negative numbers.
# Another solution is [-3,1,4,-2], which would also be accepted.
# Example 3:
#
# Input: adjacentPairs = [[100000,-100000]]
# Output: [100000,-100000]
#
#
# Constraints:
#
# nums.length == n
# adjacentPairs.length == n - 1
# adjacentPairs[i].length == 2
# 2 <= n <= 105
# -105 <= nums[i], ui, vi <= 105
# There exists some nums that has adjacentPairs as its pairs.
from collections import defaultdict
from typing import List


class Solution:
    def restoreArray(self, adjacentPairs: List[List[int]]) -> List[int]:
        g = defaultdict(list)
        for a, b in adjacentPairs:
            g[a].append(b)
            g[b].append(a)
        n = len(adjacentPairs) + 1
        ans = [0] * n
        for i, v in g.items():
            if len(v) == 1:
                ans[0] = i
                ans[1] = v[0]
                break
        for i in range(2, n):
            v = g[ans[i - 1]]
            ans[i] = v[0] if v[1] == ans[i - 2] else v[1]
        return ans


if __name__ == '__main__':
    s = Solution()
    print(s.restoreArray([[2, 1], [3, 4], [3, 2]]))
    print(s.restoreArray([[4, -2], [1, 4], [-3, 1]]))
    print(s.restoreArray([[100000, -100000]]))

#Alice plays the following game, loosely based on the card game "21".
from functools import cache


# Alice starts with 0 points and draws numbers while she has less than k points. During each draw, she gains an integer number of points randomly from the range [1, maxPts], where maxPts is an integer. Each draw is independent and the outcomes have equal probabilities.
#
# Alice stops drawing numbers when she gets k or more points.
#
# Return the probability that Alice has n or fewer points.
#
# Answers within 10-5 of the actual answer are considered accepted.
#
#
#
# Example 1:
#
# Input: n = 10, k = 1, maxPts = 10
# Output: 1.00000
# Explanation: Alice gets a single card, then stops.
# Example 2:
#
# Input: n = 6, k = 1, maxPts = 10
# Output: 0.60000
# Explanation: Alice gets a single card, then stops.
# In 6 out of 10 possibilities, she is at or below 6 points.
# Example 3:
#
# Input: n = 21, k = 17, maxPts = 10
# Output: 0.73278
#
#
# Constraints:
#
# 0 <= k <= n <= 104
# 1 <= maxPts <= 104



class Solution:
    def new21Game(self, n: int, k: int, maxPts: int) -> float:
        @cache
        def dfs(i: int) -> float:
            if i >= k:
                return int(i <= n)
            if i == k - 1:
                return min(n - k + 1, maxPts) / maxPts
            return dfs(i + 1) + (dfs(i + 1) - dfs(i + maxPts + 1)) / maxPts

        return dfs(0)

if __name__ == '__main__':
    s = Solution()
    print(s.new21Game(n=10, k=1, maxPts=10))
    print(s.new21Game(n=6, k=1, maxPts=10))
    print(s.new21Game(n=21, k=17, maxPts=10))
    print(s.new21Game(n=9811, k=8776, maxPts=1096))






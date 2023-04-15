# There are n piles of coins on a table. Each pile consists of a positive number of coins of assorted denominations.
#
# In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.
#
# Given a list piles, where piles[i] is a list of integers denoting the composition of the ith pile from top to bottom, and a positive integer k, return the maximum total value of coins you can have in your wallet if you choose exactly k coins optimally.
#
#
#
# Example 1:
#
#
# Input: piles = [[1,100,3],[7,8,9]], k = 2
# Output: 101
# Explanation:
# The above diagram shows the different ways we can choose k coins.
# The maximum total we can obtain is 101.
# Example 2:
#
# Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
# Output: 706
# Explanation:
# The maximum total can be obtained if we choose all coins from the last pile.
#
#
# Constraints:
#
# n == piles.length
# 1 <= n <= 1000
# 1 <= piles[i][j] <= 105
# 1 <= k <= sum(piles[i].length) <= 2000
import functools
from typing import List


class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        # dp(i, k) := max value of picking k coins from piles[i:]
        @functools.lru_cache(None)
        def dp(i: int, k: int) -> int:
            if i == len(piles) or k == 0:
                return 0

            ans = dp(i + 1, k)  # Pick 0 coins from current pile
            val = 0  # Coins picked from current pile

            # Try to pick 1, 2, ..., k coins from current pile
            for j in range(min(len(piles[i]), k)):
                val += piles[i][j]
                ans = max(ans, val + dp(i + 1, k - j - 1))

            return ans

        return dp(0, k)


if __name__ == '__main__':
    solver = Solution()
    cases = [
        ([[1, 100, 3], [7, 8, 9]], 2),
        ([[100], [100], [100], [100], [100], [100], [1, 1, 1, 1, 1, 1, 700]], 7),
    ]
    rslts = [solver.maxValueOfCoins(piles, k) for piles, k in cases]
    for cs, rs in zip(cases, rslts):
        print(f"case: {cs} | solution: {rs}")

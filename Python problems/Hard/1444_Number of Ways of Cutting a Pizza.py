# Given a rectangular pizza represented as a rows x cols matrix containing the following characters: 'A' (an apple) and '.' (empty cell) and given the integer k. You have to cut the pizza into k pieces using k-1 cuts.
#
# For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces. If you cut the pizza vertically, give the left part of the pizza to a person. If you cut the pizza horizontally, give the upper part of the pizza to a person. Give the last piece of pizza to the last person.
#
# Return the number of ways of cutting the pizza such that each piece contains at least one apple. Since the answer can be a huge number, return this modulo 10^9 + 7.
#
#
#
# Example 1:
#
#
#
# Input: pizza = ["A..","AAA","..."], k = 3
# Output: 3
# Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.
# Example 2:
#
# Input: pizza = ["A..","AA.","..."], k = 3
# Output: 1
# Example 3:
#
# Input: pizza = ["A..","A..","..."], k = 1
# Output: 1
#
#
# Constraints:
#
# 1 <= rows, cols <= 50
# rows == pizza.length
# cols == pizza[i].length
# 1 <= k <= 10
# pizza consists of characters 'A' and '.' only.
import functools
from typing import List


class Solution:
    def ways(self, pizza: List[str], k: int) -> int:
        M = len(pizza)
        N = len(pizza[0])
        prefix = [[0] * (N + 1) for _ in range(M + 1)]

        for i in range(M):
            for j in range(N):
                prefix[i + 1][j + 1] = (pizza[i][j] == 'A') + \
                                       prefix[i][j + 1] + prefix[i + 1][j] - prefix[i][j]

        # HasApple of pizza[row1..row2)[col1..col2)
        def hasApple(row1: int, row2: int, col1: int, col2: int) -> bool:
            return (prefix[row2][col2] - prefix[row1][col2] -
                    prefix[row2][col1] + prefix[row1][col1]) > 0

        # Dp(m, n, k) := # of ways to cut pizza[m:M][n:N] w/ k cuts
        @functools.lru_cache(None)
        def dp(m: int, n: int, k: int) -> int:
            if k == 0:
                return 1

            ans = 0

            for i in range(m + 1, M):  # Cut horizontally
                if hasApple(m, i, n, N) and hasApple(i, M, n, N):
                    ans += dp(i, n, k - 1)

            for j in range(n + 1, N):  # Cut vertically
                if hasApple(m, M, n, j) and hasApple(m, M, j, N):
                    ans += dp(m, j, k - 1)

            return ans

        return dp(0, 0, k - 1) % (10 ** 9 + 7)


if __name__ == '__main__':
    s = Solution()
    print(s.ways(["A..", "AAA", "..."], 3))
    print(s.ways(["A..", "AA.", "..."], 3))
    print(s.ways(["A..", "A..", "..."], 1))

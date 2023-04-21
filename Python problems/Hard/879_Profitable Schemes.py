# There is a group of n members, and a list of various crimes they could commit. The ith crime generates a profit[i] and requires group[i] members to participate in it. If a member participates in one crime, that member can't participate in another crime.
#
# Let's call a profitable scheme any subset of these crimes that generates at least minProfit profit, and the total number of members participating in that subset of crimes is at most n.
#
# Return the number of schemes that can be chosen. Since the answer may be very large, return it modulo 109 + 7.
#
#
#
# Example 1:
#
# Input: n = 5, minProfit = 3, group = [2,2], profit = [2,3]
# Output: 2
# Explanation: To make a profit of at least 3, the group could either commit crimes 0 and 1, or just crime 1.
# In total, there are 2 schemes.
# Example 2:
#
# Input: n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]
# Output: 7
# Explanation: To make a profit of at least 5, the group could commit any crimes, as long as they commit one.
# There are 7 possible schemes: (0), (1), (2), (0,1), (0,2), (1,2), and (0,1,2).
#
#
# Constraints:
#
# 1 <= n <= 100
# 0 <= minProfit <= 100
# 1 <= group.length <= 100
# 1 <= group[i] <= 100
# profit.length == group.length
# 0 <= profit[i] <= 100
import itertools
from typing import List


class Solution:
    def profitableSchemes(self, n: int, minProfit: int, group: List[int], profit: List[int]) -> int:
        mod = 10 ** 9 + 7  # модуль для возврата ответа
        dp = [[[0] * (minProfit + 1) for _ in range(n + 1)] for _ in range(len(group) + 1)]  # создаем трехмерный
        # массив dp размером (len(group) + 1) x (n + 1) x (minProfit + 1)

        # Нет преступлений, нет прибыли, и любое количество участников
        for i in range(n + 1):
            dp[0][i][0] = 1

        for k in range(1, len(group) + 1):
            g = group[k - 1]
            p = profit[k - 1]
            for i, j in itertools.product(range(n + 1), range(minProfit + 1)):
                dp[k][i][j] = (
                    dp[k - 1][i][j]
                    if i < g
                    else (dp[k - 1][i][j] + dp[k - 1][i - g][max(0, j - p)]) % mod
                )
        return dp[len(group)][n][minProfit]


if __name__ == '__main__':
    solver = Solution()
    cases = [
        (5, 3, [2, 2], [2, 3]),
        (10, 5, [2, 3, 5], [6, 7, 8]),
    ]
    for n, minProfit, group, profit in cases:
        print(solver.profitableSchemes(n, minProfit, group, profit))

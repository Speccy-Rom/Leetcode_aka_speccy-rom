# Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
#
# You have the following three operations permitted on a word:
#
# Insert a character
# Delete a character
# Replace a character
#
#
# Example 1:
#
# Input: word1 = "horse", word2 = "ros"
# Output: 3
# Explanation:
# horse -> rorse (replace 'h' with 'r')
# rorse -> rose (remove 'r')
# rose -> ros (remove 'e')
# Example 2:
#
# Input: word1 = "intention", word2 = "execution"
# Output: 5
# Explanation:
# intention -> inention (remove 't')
# inention -> enention (replace 'i' with 'e')
# enention -> exention (replace 'n' with 'x')
# exention -> exection (replace 'n' with 'c')
# exection -> execution (insert 'u')
#
#
# Constraints:
#
# 0 <= word1.length, word2.length <= 500
# word1 and word2 consist of lowercase English letters.


# Для решения данной задачи будем использовать динамическое программирование и алгоритм Левенштейна.
#
# Алгоритм Левенштейна - это алгоритм нахождения расстояния редактирования между двумя строками, т.е. минимальное
# количество операций, необходимых для преобразования одной строки в другую. Операции могут быть такими: вставка,
# удаление и замена символа.
#
# Для решения данной задачи мы будем использовать двумерный массив, где dp[i][j] будет хранить количество операций,
# необходимых для преобразования первых i символов word1 в первые j символов word2. Начальное значение dp[0][0] будет
# равно 0, т.к. для преобразования пустой строки в пустую строку никаких операций не требуется.
#
# Значения в массиве dp будут заполнены следующим образом:
#
# если word1[i-1] == word2[j-1], то dp[i][j] = dp[i-1][j-1] (т.к. никаких операций не требуется) если word1[i-1] !=
# word2[j-1], то dp[i][j] будет минимальным из трех возможных значений: dp[i-1][j-1] + 1 (замена), dp[i-1][j] + 1 (
# удаление) и dp[i][j-1] + 1 (вставка) Итоговый результат будет храниться в dp[len(word1)][len(word2)]. Время работы
# алгоритма - O(m*n), где m и n - длины строк word1 и word2 соответственно.

import itertools


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m = len(word1)
        n = len(word2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]

        for i, j in itertools.product(range(m + 1), range(n + 1)):
            if i == 0:
                dp[i][j] = j
            elif j == 0:
                dp[i][j] = i
            elif word1[i - 1] == word2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1]
            else:
                dp[i][j] = 1 + min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1])

        return dp[m][n]


if __name__ == '__main__':
    res = Solution()
    print(res.minDistance("horse", "ros"))
    print(res.minDistance("intention", "execution"))
    print(res.minDistance("zoologicoarchaeologist", "zoogeologist"))


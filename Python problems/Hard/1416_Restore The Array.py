# A program was supposed to print an array of integers. The program forgot to print whitespaces and the array is printed as a string of digits s and all we know is that all integers in the array were in the range [1, k] and there are no leading zeros in the array.
#
# Given the string s and the integer k, return the number of the possible arrays that can be printed as s using the mentioned program. Since the answer may be very large, return it modulo 109 + 7.
#
#
#
# Example 1:
#
# Input: s = "1000", k = 10000
# Output: 1
# Explanation: The only possible array is [1000]
# Example 2:
#
# Input: s = "1000", k = 10
# Output: 0
# Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.
# Example 3:
#
# Input: s = "1317", k = 2000
# Output: 8
# Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
#
#
# Constraints:
#
# 1 <= s.length <= 105
# s consists of only digits and does not contain leading zeros.
# 1 <= k <= 109


class Solution:
    def numOfArrays(self, s: str, k: int) -> int:
        mod = 1000000007
        n = len(s)
        dp = [0] * (n + 1)
        dp[0] = 1

        for i in range(1, n + 1):
            for j in range(i - 1, -1, -1):
                if i - j > 10:
                    break
                if s[j] == '0':
                    continue
                num = int(s[j:i])
                if num <= k:
                    dp[i] = (dp[i] + dp[j]) % mod

        return dp[n]


if __name__ == '__main__':
    solver = Solution()
    cases = [
        ("1000", 10000),
        ("1000", 10),
        ("1317", 2000),
    ]
    for s, k in cases:
        print(solver.numOfArrays(s, k))


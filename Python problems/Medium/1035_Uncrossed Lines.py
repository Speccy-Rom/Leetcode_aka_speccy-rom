import itertools
class Solution:
    def maxUncrossedLines(self, nums1: List[int], nums2: List[int]) -> int:
        m, n = len(nums1), len(nums2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i, j in itertools.product(range(1, m + 1), range(1, n + 1)):
            dp[i][j] = (
                dp[i - 1][j - 1] + 1
                if nums1[i - 1] == nums2[j - 1]
                else max(dp[i - 1][j], dp[i][j - 1])
            )
        return dp[m][n]


if __name__ == '__main__':
    s = Solution()
    print(s.maxUncrossedLines([1, 4, 2], [1, 2, 4]))
    print(s.maxUncrossedLines([2, 5, 1, 2, 5], [10, 5, 2, 1, 5, 2]))
    print(s.maxUncrossedLines([1, 3, 7, 1, 7, 5], [1, 9, 2, 5, 1]))
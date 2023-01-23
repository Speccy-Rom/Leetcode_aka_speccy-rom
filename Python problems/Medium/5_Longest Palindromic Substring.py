# Given a string s, return the longest palindromic substring in s.
#
# Example 1:
#
# Input: s = "babad"
# Output: "bab"
# Explanation: "aba" is also a valid answer.
# Example 2:
#
# Input: s = "cbbd"
# Output: "bb"
#
#
# Constraints:
#
# 1 <= s.length <= 1000
# s consist of only digits and English letters.

class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) < 2 or s == s[::-1]:
            return s

        start = 0
        end = 0
        for i in range(len(s)):
            len1 = self.expand(s, i, i)
            len2 = self.expand(s, i, i + 1)
            length = max(len1, len2)
            if length > end - start:
                start = i - (length - 1) // 2
                end = i + length // 2
        return s[start:end + 1]

    def expand(self, s: str, left: int, right: int) -> int:
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return right - left - 1


# class Solution:
#   def longestPalindrome(self, s: str) -> str:
#     if not s:
#       return ''
#
#     indices = [0, 0]
#
#     # Returns [start, end] indices of the longest palindrome extended from s[i..j]
#     def extend(s: str, i: int, j: int) -> Tuple[int, int]:
#       while i >= 0 and j < len(s):
#         if s[i] != s[j]:
#           break
#         i -= 1
#         j += 1
#       return i + 1, j - 1
#
#     for i in range(len(s)):
#       l1, r1 = extend(s, i, i)
#       if r1 - l1 > indices[1] - indices[0]:
#         indices = l1, r1
#       if i + 1 < len(s) and s[i] == s[i + 1]:
#         l2, r2 = extend(s, i, i + 1)
#         if r2 - l2 > indices[1] - indices[0]:
#           indices = l2, r2
#
#     return s[indices[0]:indices[1] + 1]

if __name__ == "__main__":
    solution = Solution()
    print(solution.longestPalindrome("babad"))
    print(solution.longestPalindrome("cbbd"))
    print(solution.longestPalindrome("a"))
    print(solution.longestPalindrome("ac"))
    print(solution.longestPalindrome("bb"))

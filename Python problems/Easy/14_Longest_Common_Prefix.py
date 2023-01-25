# Write a function to find the longest common prefix string amongst an array of strings.
# If there is no common prefix, return an empty string "".
#
# Example 1:
#
# Input: strs = ["flower","flow","flight"]
# Output: "fl"
# Example 2:
#
# Input: strs = ["dog","racecar","car"]
# Output: ""
# Explanation: There is no common prefix among the input strings.
#
#
# Constraints:
#
# 1 <= strs.length <= 200
# 0 <= strs[i].length <= 200
# strs[i] consists of only lowercase English letters.
from typing import List


# class Solution:
#     def longestCommonPrefix(self, strs: List[str]) -> str:
#         if not strs:
#             return ""
#         prefix = strs[0]
#         for i in range(1, len(strs)):
#             while strs[i].find(prefix) != 0:
#                 prefix = prefix[:-1]
#                 if not prefix:
#                     return ""
#         return prefix

# class Solution:
#   def longestCommonPrefix(self, strs: List[str]) -> str:
#     if not strs:
#       return ''
#
#     for i in range(len(strs[0])):
#       for j in range(1, len(strs)):
#         if i == len(strs[j]) or strs[j][i] != strs[0][i]:
#           return strs[0][:i]
#
#     return strs[0]

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        pre = strs[0]
        for i in strs:
            while not i.startswith(pre):
                pre = pre[:-1]
        return pre


if __name__ == "__main__":
    solution = Solution()
    print(solution.longestCommonPrefix(["flower", "flow", "flight"]))
    print(solution.longestCommonPrefix(["dog", "racecar", "car"]))

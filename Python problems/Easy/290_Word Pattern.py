# Given a pattern and a string s, find if s follows the same pattern.
#
# Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
#
#
#
# Example 1:
#
# Input: pattern = "abba", s = "dog cat cat dog"
# Output: true
# Example 2:
#
# Input: pattern = "abba", s = "dog cat cat fish"
# Output: false
# Example 3:
#
# Input: pattern = "aaaa", s = "dog cat cat dog"
# Output: false


# class Solution:
#     def wordPattern(self, pattern: str, s: str) -> bool:
#         words = s.split()
#         return list(map(pattern.find, pattern)) == list(map(words.index, words))

# class Solution:
#     def wordPattern(self, pattern: str, s: str) -> bool:
#         s = s.split()
#         return len(pattern) == len(s) and len(set(pattern)) == len(set(s)) == len(set(zip(pattern, s)))

class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        obj = {}
        objPattern = {}

        if len(pattern) != len(s.split(' ')):
            return False
        return len(set(s.split(' '))) == len(set(pattern)) and len(set(s.split(' '))) == len(
            set(zip(pattern, s.split(' '))))


if __name__ == "__main__":
    solution = Solution()
    print(solution.wordPattern("abba", "dog cat cat dog"))
    print(solution.wordPattern("abba", "dog cat cat fish"))
    print(solution.wordPattern("aaaa", "dog cat cat dog"))

# Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).
#
#
#
# Example 1:
#
# Input: low = 3, high = 7
# Output: 3
# Explanation: The odd numbers between 3 and 7 are [3,5,7].
# Example 2:
#
# Input: low = 8, high = 10
# Output: 1
# Explanation: The odd numbers between 8 and 10 are [9].
#
#
# Constraints:
#
# 0 <= low <= high <= 10^9


class Solution:
    def countOdds(self, low: int, high: int) -> int:
        return (high - low) // 2 + (low % 2 or high % 2)


# class Solution:
#     def countOdds(self, low: int, high: int) -> int:
#         return (high + 1) // 2 - low // 2


if __name__ == '__main__':
    s = Solution()
    print(s.countOdds(3, 7))  # 3
    print(s.countOdds(8, 10))  # 1
    print(s.countOdds(0, 10))  # 6
    print(s.countOdds(1, 10))  # 5
    print(s.countOdds(2, 10))  # 4
    print(s.countOdds(3, 10))  # 4
    print(s.countOdds(4, 10))  # 3
    print(s.countOdds(5, 10))  # 3
    print(s.countOdds(6, 10))  # 2
    print(s.countOdds(7, 10))  # 2
    print(s.countOdds(8, 10))  # 1
    print(s.countOdds(9, 10))  # 1
    print(s.countOdds(10, 10))  # 1

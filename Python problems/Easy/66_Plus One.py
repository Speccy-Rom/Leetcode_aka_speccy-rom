# You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
#
# Increment the large integer by one and return the resulting array of digits.
#
#
#
# Example 1:
#
# Input: digits = [1,2,3]
# Output: [1,2,4]
# Explanation: The array represents the integer 123.
# Incrementing by one gives 123 + 1 = 124.
# Thus, the result should be [1,2,4].
# Example 2:
#
# Input: digits = [4,3,2,1]
# Output: [4,3,2,2]
# Explanation: The array represents the integer 4321.
# Incrementing by one gives 4321 + 1 = 4322.
# Thus, the result should be [4,3,2,2].
# Example 3:
#
# Input: digits = [9]
# Output: [1,0]
# Explanation: The array represents the integer 9.
# Incrementing by one gives 9 + 1 = 10.
# Thus, the result should be [1,0].
from typing import List


class Solution:
    def plus_one(self, digits: List[int]) -> List[int]:
        digits[-1] += 1
        for i in range(len(digits) - 1, -1, -1):
            if digits[i] > 9:
                digits[i] = 0
                if i == 0:
                    digits.insert(0, 1)
                else:
                    digits[i - 1] += 1
        return digits


if __name__ == "__main__":
    solution = Solution()
    print(solution.plus_one([9]))
    print(solution.plus_one([1, 2, 3]))
    print(solution.plus_one([4, 3, 2, 1]))
    print(solution.plus_one([9, 9, 9, 9]))
    print(solution.plus_one([1, 9, 9, 9]))
    print(solution.plus_one([9, 9, 9, 1]))
    print(solution.plus_one([9, 9, 1, 9]))
    print(solution.plus_one([9, 1, 9, 9]))
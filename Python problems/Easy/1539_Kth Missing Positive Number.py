# Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
#
# Return the kth positive integer that is missing from this array.
#
#
#
# Example 1:
#
# Input: arr = [2,3,4,7,11], k = 5
# Output: 9
# Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
# Example 2:
#
# Input: arr = [1,2,3,4], k = 2
# Output: 6
# Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
#
#
# Constraints:
#
# 1 <= arr.length <= 1000
# 1 <= arr[i] <= 1000
# 1 <= k <= 1000
# arr[i] < arr[j] for 1 <= i < j <= arr.length
#
#
# Follow up:
#
# Could you solve this problem in less than O(n) complexity?
from typing import List


class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        i = 0
        for num in range(1, arr[-1] + k + 1):
            if i < len(arr) and num == arr[i]:
                i += 1
            else:
                k -= 1
            if k == 0:
                return num


if __name__ == '__main__':
    s = Solution()
    print(s.findKthPositive([2, 3, 4, 7, 11], 5))
    print(s.findKthPositive([1, 2, 3, 4], 2))

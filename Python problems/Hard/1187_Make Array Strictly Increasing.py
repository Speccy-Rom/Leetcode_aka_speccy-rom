#Given two integer arrays arr1 and arr2, return the minimum number of operations (possibly zero) needed to make arr1 strictly increasing.
from bisect import bisect_left
from math import inf
from typing import List


# In one operation, you can choose two indices 0 <= i < arr1.length and 0 <= j < arr2.length and do the assignment arr1[i] = arr2[j].
#
# If there is no way to make arr1 strictly increasing, return -1.
#
#
#
# Example 1:
#
# Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
# Output: 1
# Explanation: Replace 5 with 2, then arr1 = [1, 2, 3, 6, 7].
# Example 2:
#
# Input: arr1 = [1,5,3,6,7], arr2 = [4,3,1]
# Output: 2
# Explanation: Replace 5 with 3 and then replace 3 with 4. arr1 = [1, 3, 4, 6, 7].
# Example 3:
#
# Input: arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
# Output: -1
# Explanation: You can't make arr1 strictly increasing.
#
#
# Constraints:
#
# 1 <= arr1.length, arr2.length <= 2000
# 0 <= arr1[i], arr2[i] <= 10^9


class Solution:
    def makeArrayIncreasing(self, arr1: List[int], arr2: List[int]) -> int:
        arr2.sort()
        m = 0
        for x in arr2:
            if m == 0 or x != arr2[m - 1]:
                arr2[m] = x
                m += 1
        arr2 = arr2[:m]
        arr = [-inf] + arr1 + [inf]
        n = len(arr)
        f = [inf] * n
        f[0] = 0
        for i in range(1, n):
            if arr[i - 1] < arr[i]:
                f[i] = f[i - 1]
            j = bisect_left(arr2, arr[i])
            for k in range(1, min(i - 1, j) + 1):
                if arr[i - k - 1] < arr2[j - k]:
                    f[i] = min(f[i], f[i - k - 1] + k)
        return -1 if f[n - 1] >= inf else f[n - 1]



if __name__ == '__main__':
    solution = Solution()
    print(solution.makeArrayIncreasing([1, 5, 3, 6, 7], [1, 3, 2, 4]))
    print(solution.makeArrayIncreasing([1, 5, 3, 6, 7], [4, 3, 1]))
    print(solution.makeArrayIncreasing([1, 5, 3, 6, 7], [1, 6, 3, 3]))
# Given a non-negative integer x, return the square root of x rounded
# down to the nearest integer. The returned integer should be non-negative as well.
#
# You must not use any built-in exponent function or operator.
# For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
#
#
# Example 1:
#
# Input: x = 4
# Output: 2
# Explanation: The square root of 4 is 2, so we return 2.
# Example 2:
#
# Input: x = 8
# Output: 2
# Explanation: The square root of 8 is 2.82842..., and since we round it down to the
# nearest integer, 2 is returned.

class Solution:
    def mySqrt(self, x: int) -> int:
        # return int(x ** 0.5)
        left = 0
        right = x
        while left <= right:
            mid = (left + right) // 2
            if mid * mid == x:
                return mid
            elif mid * mid < x:
                left = mid + 1
            else:
                right = mid - 1
        return right


if __name__ == "__main__":
    mySqrt = Solution()
    print(mySqrt.mySqrt(4))
    print(mySqrt.mySqrt(8))
    print(mySqrt.mySqrt(9))
    print(mySqrt.mySqrt(10))
    print(mySqrt.mySqrt(11))
    print(mySqrt.mySqrt(12))


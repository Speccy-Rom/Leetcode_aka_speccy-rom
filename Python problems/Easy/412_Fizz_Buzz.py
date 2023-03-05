# Given an integer n, return a string array answer (1-indexed) where:
#
# answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
# answer[i] == "Fizz" if i is divisible by 3.
# answer[i] == "Buzz" if i is divisible by 5.
# answer[i] == i (as a string) if none of the above conditions are true.
#
#
# Example 1:
#
# Input: n = 3
# Output: ["1","2","Fizz"]
# Example 2:
#
# Input: n = 5
# Output: ["1","2","Fizz","4","Buzz"]
# Example 3:
#
# Input: n = 15
# Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
#
#
# Constraints:
#
# 1 <= n <= 104
from typing import List


class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        return ["Fizz" * (i % 3 == 0) + "Buzz" * (i % 5 == 0) or str(i) for i in range(1, n + 1)]


# class Solution:
#     def fizzBuzz(self, n: int) -> List[str]:
#         d = {3: 'Fizz', 5: 'Buzz'}
#         return [''.join([d[k] for k in d if i % k == 0]) or str(i) for i in range(1, n + 1)]


if __name__ == "__main__":
    solution = Solution()
    print(solution.fizzBuzz(15))

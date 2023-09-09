import itertools
from typing import List


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        f = [1] + [0] * target
        for i, x in itertools.product(range(1, target + 1), nums):
            if i >= x:
                f[i] += f[i - x]
        return f[target]


if __name__ == '__main__':

    def test(input1, input2):
        Test = Solution()
        ans = Test.combinationSum4(input1, input2)
        print(ans)
        return ans

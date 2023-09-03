import itertools


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        f = [1] * n
        for _, j in itertools.product(range(1, m), range(1, n)):
            f[j] += f[j - 1]
        return f[-1]


if __name__ == '__main__':
    print(Solution().uniquePaths(3, 7))
    print(Solution().uniquePaths(3, 2))
    print(Solution().uniquePaths(7, 3))
    print(Solution().uniquePaths(3, 3))

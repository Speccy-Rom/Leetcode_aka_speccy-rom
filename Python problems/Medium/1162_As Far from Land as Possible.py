import itertools
from typing import List


class Solution:
    def maxDistance(self, grid: List[List[int]]) -> int:
        n = len(grid)
        queue = [
            (i, j)
            for i, j in itertools.product(range(n), range(n))
            if grid[i][j] == 1
        ]
        if len(queue) in [0, n ** 2]:
            return -1
        distance = -1
        while queue:
            distance += 1
            for _ in range(len(queue)):
                i, j = queue.pop(0)
                for x, y in [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]:
                    if 0 <= x < n and 0 <= y < n and grid[x][y] == 0:
                        grid[x][y] = 1
                        queue.append((x, y))
        return distance


if __name__ == '__main__':
    s = Solution()
    print(s.maxDistance([[1, 0, 1], [0, 0, 0], [1, 0, 1]]))
    print(s.maxDistance([[1, 0, 0], [0, 0, 0], [0, 0, 0]]))

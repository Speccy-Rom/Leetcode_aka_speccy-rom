# Given an m x n binary matrix mat, return the number of special positions in mat.

# A position (i, j) is called special if mat[i][j] == 1 and all other elements in row i and column j are 0 (rows and columns are 0-indexed).
#
#
#
# Example 1:
#
#
# Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
# Output: 1
# Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
# Example 2:
#
#
# Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
# Output: 3
# Explanation: (0, 0), (1, 1) and (2, 2) are special positions.
#
#
# Constraints:
#
# m == mat.length
# n == mat[i].length
# 1 <= m, n <= 100
# mat[i][j] is either 0 or 1.


class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        r = [0] * m
        c = [0] * n
        for i, row in enumerate(mat):
            for j, v in enumerate(row):
                r[i] += v
                c[j] += v
        ans = 0
        for i in range(m):
            for j in range(n):
                if mat[i][j] == 1 and r[i] == 1 and c[j] == 1:
                    ans += 1
        return ans


if __name__ == '__main__':
    s = Solution()
    print(s.numSpecial([[1, 0, 0], [0, 0, 1], [1, 0, 0]]))
    print(s.numSpecial([[1, 0, 0], [0, 1, 0], [0, 0, 1]]))
    print(s.numSpecial([[0, 0, 0, 1], [1, 0, 0, 0], [0, 1, 1, 0], [0, 0, 0, 0]]))

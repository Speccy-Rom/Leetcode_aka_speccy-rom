/*
You are given an m x n integer matrix grid, where you can move from a cell to any adjacent cell in all 4 directions.

Return the number of strictly increasing paths in the grid such that you can start from any cell and end at any cell. Since the answer may be very large, return it modulo 109 + 7.

Two paths are considered different if they do not have exactly the same sequence of visited cells.



Example 1:


Input: grid = [[1,1],[3,4]]
Output: 8
Explanation: The strictly increasing paths are:
- Paths with length 1: [1], [1], [3], [4].
- Paths with length 2: [1 -> 3], [1 -> 4], [3 -> 4].
- Paths with length 3: [1 -> 3 -> 4].
The total number of paths is 4 + 3 + 1 = 8.
Example 2:

Input: grid = [[1],[2]]
Output: 3
Explanation: The strictly increasing paths are:
- Paths with length 1: [1], [2].
- Paths with length 2: [1 -> 2].
The total number of paths is 2 + 1 = 3.


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 1000
1 <= m * n <= 105
1 <= grid[i][j] <= 105
*/

package main

import "fmt"

func countPaths(grid [][]int) (ans int) {
	const mod = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = 1
		dirs := [5]int{-1, 0, 1, 0, -1}
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[i][j] < grid[x][y] {
				f[i][j] = (f[i][j] + dfs(x, y)) % mod
			}
		}
		return f[i][j]
	}
	for i, row := range grid {
		for j := range row {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return
}

func main() {
	fmt.Println(countPaths([][]int{{1, 1}, {3, 4}}))
	fmt.Println(countPaths([][]int{{1}, {2}}))

}

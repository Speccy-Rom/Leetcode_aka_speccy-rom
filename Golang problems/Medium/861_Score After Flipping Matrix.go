/*
You are given an m x n binary matrix grid.

A move consists of choosing any row or column and toggling each value in that row or column (i.e., changing all 0's to 1's, and all 1's to 0's).

Every row of the matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.

Return the highest possible score after making any number of moves (including zero moves).



Example 1:


Input: grid = [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
Output: 39
Explanation: 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
Example 2:

Input: grid = [[0]]
Output: 1


*/

package main

import "fmt"

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}
	ans := 0
	for j := 0; j < n; j++ {
		cnt := 0
		for i := 0; i < m; i++ {
			cnt += grid[i][j]
		}
		if cnt < m-cnt {
			cnt = m - cnt
		}
		ans += cnt * (1 << (n - j - 1))
	}
	return ans
}

// Test cases
func main() {
	// Test cases
	res := matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}})
	fmt.Println(res)
	res = matrixScore([][]int{{0}})
	fmt.Println(res)
}

/*
Given a rectangular pizza represented as a rows x cols matrix containing the following characters: 'A' (an apple) and '.' (empty cell) and given the integer k. You have to cut the pizza into k pieces using k-1 cuts.

For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces. If you cut the pizza vertically, give the left part of the pizza to a person. If you cut the pizza horizontally, give the upper part of the pizza to a person. Give the last piece of pizza to the last person.

Return the number of ways of cutting the pizza such that each piece contains at least one apple. Since the answer can be a huge number, return this modulo 10^9 + 7.



Example 1:



Input: pizza = ["A..","AAA","..."], k = 3
Output: 3
Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.
Example 2:

Input: pizza = ["A..","AA.","..."], k = 3
Output: 1
Example 3:

Input: pizza = ["A..","A..","..."], k = 1
Output: 1


Constraints:

1 <= rows, cols <= 50
rows == pizza.length
cols == pizza[i].length
1 <= k <= 10
pizza consists of characters 'A' and '.' only.
*/

package main

import "fmt"

func ways(pizza []string, k int) int {
	m, n := len(pizza), len(pizza[0])
	prefix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		prefix[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prefix[i][j] = prefix[i+1][j] + prefix[i][j+1] - prefix[i+1][j+1]
			if pizza[i][j] == 'A' {
				prefix[i][j]++
			}
		}

	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

	}
	return dfs(pizza, prefix, dp, 0, 0, k)

}

func dfs(pizza []string, prefix [][]int, dp [][]int, i, j, k int) int {
	if prefix[i][j] == 0 {
		return 0
	}
	if k == 1 {
		return 1
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	m, n := len(pizza), len(pizza[0])
	for x := i + 1; x < m; x++ {
		if prefix[i][j]-prefix[x][j] > 0 {
			dp[i][j] = (dp[i][j] + dfs(pizza, prefix, dp, x, j, k-1)) % 1000000007
		}
	}
	for y := j + 1; y < n; y++ {
		if prefix[i][j]-prefix[i][y] > 0 {
			dp[i][j] = (dp[i][j] + dfs(pizza, prefix, dp, i, y, k-1)) % 1000000007
		}
	}
	return dp[i][j]
}

func main() {
	pizza := []string{"A..", "AAA", "..."}
	k := 3
	fmt.Println(ways(pizza, k))
}

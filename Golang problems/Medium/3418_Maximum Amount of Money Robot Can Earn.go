package main

import "math"

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	f := make([][][]int, m)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, 3)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}

	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i >= m || j >= n {
			return math.MinInt32 / 2
		}
		if f[i][j][k] != math.MinInt32 {
			return f[i][j][k]
		}
		if i == m-1 && j == n-1 {
			if k > 0 {
				return maxInt(0, coins[i][j])
			}
			return coins[i][j]
		}

		ans := coins[i][j] + maxInt(dfs(i+1, j, k), dfs(i, j+1, k))
		if coins[i][j] < 0 && k > 0 {
			ans = maxInt(ans, maxInt(dfs(i+1, j, k-1), dfs(i, j+1, k-1)))
		}
		f[i][j][k] = ans
		return ans
	}

	return dfs(0, 0, 2)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumAmount([][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}})
}

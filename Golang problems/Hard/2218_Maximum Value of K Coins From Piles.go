/*
There are n piles of coins on a table. Each pile consists of a positive number of coins of assorted denominations.

In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.

Given a list piles, where piles[i] is a list of integers denoting the composition of the ith pile from top to bottom, and a positive integer k, return the maximum total value of coins you can have in your wallet if you choose exactly k coins optimally.



Example 1:


Input: piles = [[1,100,3],[7,8,9]], k = 2
Output: 101
Explanation:
The above diagram shows the different ways we can choose k coins.
The maximum total we can obtain is 101.
Example 2:

Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
Output: 706
Explanation:
The maximum total can be obtained if we choose all coins from the last pile.


Constraints:

n == piles.length
1 <= n <= 1000
1 <= piles[i][j] <= 105
1 <= k <= sum(piles[i].length) <= 2000 */

package main

import "fmt"

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	cache := make(map[[2]int]int)
	return dfs(piles, cache, 0, n-1, k)
}

func dfs(piles [][]int, cache map[[2]int]int, i, j, k int) int {
	if k == 0 {
		return 0
	}
	if v, ok := cache[[2]int{i, j}]; ok {
		return v
	}
	res := max(piles[i][0]+dfs(piles, cache, i+1, j, k-1), piles[j][len(piles[j])-1]+dfs(piles, cache, i, j-1, k-1))
	cache[[2]int{i, j}] = res
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func main() {
	fmt.Println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
}

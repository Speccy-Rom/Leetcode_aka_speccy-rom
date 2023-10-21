/*

Given an integer array nums and an integer k, return the maximum sum of a non-empty subsequence of that array such that for every two consecutive integers in the subsequence, nums[i] and nums[j], where i < j, the condition j - i <= k is satisfied.

A subsequence of an array is obtained by deleting some number of elements (can be zero) from the array, leaving the remaining elements in their original order.



Example 1:

Input: nums = [10,2,-10,5,20], k = 2
Output: 37
Explanation: The subsequence is [10, 2, 5, 20].
Example 2:

Input: nums = [-1,-2,-3], k = 1
Output: -1
Explanation: The subsequence must be non-empty, so we choose the largest number.
Example 3:

Input: nums = [10,-2,-10,-5,20], k = 2
Output: 23
Explanation: The subsequence is [10, -2, -5, 20].


Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104

*/

package main

import "math"

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := math.MinInt32
	q := []int{}
	for i, v := range nums {
		if len(q) > 0 && i-q[0] > k {
			q = q[1:]
		}
		dp[i] = v
		if len(q) > 0 && dp[q[0]] > 0 {
			dp[i] += dp[q[0]]
		}
		for len(q) > 0 && dp[q[len(q)-1]] < dp[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 2, -10, 5, 20}
	k := 2
	println(constrainedSubsetSum(nums, k))
	nums = []int{-1, -2, -3}
	k = 1
	println(constrainedSubsetSum(nums, k))
	nums = []int{10, -2, -10, -5, 20}
	k = 2
	println(constrainedSubsetSum(nums, k))
}

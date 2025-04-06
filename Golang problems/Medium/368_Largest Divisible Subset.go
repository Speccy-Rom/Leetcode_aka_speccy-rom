/*
Given a set of distinct positive integers nums, return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:

answer[i] % answer[j] == 0, or
answer[j] % answer[i] == 0
If there are multiple solutions, return any of them.



Example 1:

Input: nums = [1,2,3]
Output: [1,2]
Explanation: [1,3] is also accepted.
Example 2:

Input: nums = [1,2,4,8]
Output: [1,2,4,8]


Constraints:

1 <= nums.length <= 1000
1 <= nums[i] <= 2 * 109
All the integers in nums are unique.


*/

package main

import (
	"sort"
)

// largestDivisibleSubset finds the largest subset where every pair of elements
// satisfies the divisibility condition
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// Sort the array to ensure we only need to check one direction of divisibility
	sort.Ints(nums)

	n := len(nums)
	// dp[i] stores the size of the largest divisible subset ending at index i
	dp := make([]int, n)
	// prev[i] stores the index of the previous element in the subset
	prev := make([]int, n)

	// Initialize dp and prev arrays
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	maxSize := 1
	maxIndex := 0

	// For each number, find the largest subset it can be part of
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
				if dp[i] > maxSize {
					maxSize = dp[i]
					maxIndex = i
				}
			}
		}
	}

	// Reconstruct the subset
	result := make([]int, 0, maxSize)
	for maxIndex != -1 {
		result = append([]int{nums[maxIndex]}, result...)
		maxIndex = prev[maxIndex]
	}

	return result
}

/*
You are given an integer array nums and two integers minK and maxK.

A fixed-bound subarray of nums is a subarray that satisfies the following conditions:

The minimum value in the subarray is equal to minK.
The maximum value in the subarray is equal to maxK.
Return the number of fixed-bound subarrays.

A subarray is a contiguous part of an array.



Example 1:

Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
Output: 2
Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
Example 2:

Input: nums = [1,1,1,1], minK = 1, maxK = 1
Output: 10
Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.


Constraints:

2 <= nums.length <= 105
1 <= nums[i], minK, maxK <= 106

*/

package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	ans := 0
	j := -1
	prevMinKIndex := -1
	prevMaxKIndex := -1

	for i, num := range nums {
		if num < minK || num > maxK {
			j = i
		}
		if num == minK {
			prevMinKIndex = i
		}
		if num == maxK {
			prevMaxKIndex = i
		}
		ans += max(0, min(prevMinKIndex, prevMaxKIndex)-j)
	}

	return int64(ans)
}

func main() {
	nums := []int{1, 3, 5, 2, 7, 5}
	minK := 1
	maxK := 5
	countSubarrays(nums, minK, maxK)
}

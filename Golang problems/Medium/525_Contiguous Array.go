/*

Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.



Example 1:

Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
Example 2:

Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.


Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.

*/

package main

import "fmt"

func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	s, ans := 0, 0
	for i, v := range nums {
		if v == 0 {
			v = -1
		}
		s += v
		if j, ok := mp[s]; ok {
			ans = max(ans, i-j)
		} else {
			mp[s] = i
		}
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
	// Testing the function
	// Test case 1
	// Input: nums = [0,1]
	// Output: 2
	// Expected output: 2
	nums := []int{0, 1}
	fmt.Println(findMaxLength(nums))
	// Test case 2
	// Input: nums = [0,1,0]
	// Output: 2
	// Expected output: 2
	nums = []int{0, 1, 0}
	fmt.Println(findMaxLength(nums))
}

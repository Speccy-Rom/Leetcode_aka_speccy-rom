/*

Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104

*/

package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
				ans = max(ans, f[i])
			}
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
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

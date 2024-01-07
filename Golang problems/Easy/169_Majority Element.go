/*

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

package main

func majorityElement(nums []int) int {
	var cnt, m int
	for _, x := range nums {
		if cnt == 0 {
			m, cnt = x, 1
		} else {
			if m == x {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return m
}

func main() {
	println(majorityElement([]int{3, 2, 3}))
	println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

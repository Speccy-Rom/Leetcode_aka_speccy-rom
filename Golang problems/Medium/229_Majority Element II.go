/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.



Example 1:

Input: nums = [3,2,3]
Output: [3]
Example 2:

Input: nums = [1]
Output: [1]
Example 3:

Input: nums = [1,2]
Output: [1,2]


Constraints:

1 <= nums.length <= 5 * 104
-109 <= nums[i] <= 109


Follow up: Could you solve the problem in linear time and in O(1) space?

*/

package main

import "fmt"

func majorityElement(nums []int) []int {
	var n1, n2 int
	m1, m2 := 0, 1
	for _, m := range nums {
		if m == m1 {
			n1++
		} else if m == m2 {
			n2++
		} else if n1 == 0 {
			m1, n1 = m, 1
		} else if n2 == 0 {
			m2, n2 = m, 1
		} else {
			n1, n2 = n1-1, n2-1
		}
	}
	n1, n2 = 0, 0
	for _, m := range nums {
		if m == m1 {
			n1++
		} else if m == m2 {
			n2++
		}
	}
	var ans []int
	if n1 > len(nums)/3 {
		ans = append(ans, m1)
	}
	if n2 > len(nums)/3 {
		ans = append(ans, m2)
	}
	return ans
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3})) // [3]
	fmt.Println(majorityElement([]int{1}))       // [1]
	fmt.Println(majorityElement([]int{1, 2}))    // [1,2]
}

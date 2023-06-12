/*
You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"


Constraints:

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
All the values of nums are unique.
nums is sorted in ascending order.
*/

package main

import "strconv"

func summaryRanges(nums []int) (ans []string) {
	f := func(i, j int) string {
		if i == j {
			return strconv.Itoa(nums[i])
		}
		return strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j])
	}
	for i, j, n := 0, 0, len(nums); i < n; i = j + 1 {
		j = i
		for j+1 < n && nums[j+1] == nums[j]+1 {
			j++
		}
		ans = append(ans, f(i, j))
	}
	return
}

func main() {
	// Test case 1
	// Expected Output: ["0->2","4->5","7"]
	println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))

	// Test case 2
	// Expected Output: ["0","2->4","6","8->9"]
	println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

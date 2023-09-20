/*
You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.

Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.



Example 1:

Input: nums = [1,1,4,2,3], x = 5
Output: 2
Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
Example 2:

Input: nums = [5,6,7,8,9], x = 4
Output: -1
Example 3:

Input: nums = [3,2,20,1,1,3], x = 10
Output: 5
Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 104
1 <= x <= 109
*/

package main

func minOperations(nums []int, x int) int {
	x = -x
	for _, v := range nums {
		x += v
	}
	vis := map[int]int{0: -1}
	ans := 1 << 30
	s, n := 0, len(nums)
	for i, v := range nums {
		s += v
		if _, ok := vis[s]; !ok {
			vis[s] = i
		}
		if j, ok := vis[s-x]; ok {
			ans = min(ans, n-(i-j))
		}
	}
	if ans == 1<<30 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minOperations([]int{1, 1, 4, 2, 3}, 5)
	minOperations([]int{5, 6, 7, 8, 9}, 4)
	minOperations([]int{3, 2, 20, 1, 1, 3}, 10)
}

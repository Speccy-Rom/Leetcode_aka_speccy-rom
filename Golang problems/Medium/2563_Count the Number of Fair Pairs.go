/*

Given a 0-indexed integer array nums of size n and two integers lower and upper, return the number of fair pairs.

A pair (i, j) is fair if:

0 <= i < j < n, and
lower <= nums[i] + nums[j] <= upper


Example 1:

Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
Output: 6
Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).
Example 2:

Input: nums = [1,7,9,2,5], lower = 11, upper = 11
Output: 1
Explanation: There is a single fair pair: (2,3).
*/

package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) (ans int64) {
	sort.Ints(nums)
	for i, x := range nums {
		j := sort.Search(len(nums), func(h int) bool { return h > i && nums[h] >= lower-x })
		k := sort.Search(len(nums), func(h int) bool { return h > i && nums[h] >= upper-x+1 })
		ans += int64(k - j)
	}
	return
}

func main() {
	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)) // 6
	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))  // 1
}

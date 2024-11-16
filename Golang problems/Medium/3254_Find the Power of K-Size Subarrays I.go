/*

You are given an array of integers nums of length n and a positive integer k.

The power of an array is defined as:

Its maximum element if all of its elements are consecutive and sorted in ascending order.
-1 otherwise.
You need to find the power of all
subarrays
 of nums of size k.

Return an integer array results of size n - k + 1, where results[i] is the power of nums[i..(i + k - 1)].



Example 1:

Input: nums = [1,2,3,4,3,2,5], k = 3

Output: [3,4,-1,-1,-1]

Explanation:

There are 5 subarrays of nums of size 3:

[1, 2, 3] with the maximum element 3.
[2, 3, 4] with the maximum element 4.
[3, 4, 3] whose elements are not consecutive.
[4, 3, 2] whose elements are not sorted.
[3, 2, 5] whose elements are not consecutive.
Example 2:

Input: nums = [2,2,2,2,2], k = 4

Output: [-1,-1]

Example 3:

Input: nums = [3,2,3,2,3,2], k = 2

Output: [-1,3,-1,3,-1]



Constraints:

1 <= n == nums.length <= 500
1 <= nums[i] <= 105
1 <= k <= n
*/

package main

import (
	"fmt"
	"reflect"
)

func resultsArray(nums []int, k int) (ans []int) {
	n := len(nums)
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}
	}
	for i := k - 1; i < n; i++ {
		if f[i] >= k {
			ans = append(ans, nums[i])
		} else {
			ans = append(ans, -1)
		}
	}
	return
}

func main() {
	// Test cases
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 3, 2, 5}, 3, []int{3, 4, -1, -1, -1}},
		{[]int{2, 2, 2, 2, 2}, 4, []int{-1, -1}},
		{[]int{3, 2, 3, 2, 3, 2}, 2, []int{-1, 3, -1, 3, -1}},
	}
	for _, t := range tests {
		if got := resultsArray(t.nums, t.k); !reflect.DeepEqual(got, t.want) {
			fmt.Printf("resultsArray(%v, %v) = %v; want %v\n", t.nums, t.k, got, t.want)
		}
	}

}

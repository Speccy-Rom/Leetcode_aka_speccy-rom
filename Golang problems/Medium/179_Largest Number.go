/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.



Example 1:

Input: nums = [10,2]
Output: "210"
Example 2:

Input: nums = [3,30,34,5,9]
Output: "9534330"


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/

package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	vs := make([]string, len(nums))
	for i, v := range nums {
		vs[i] = strconv.Itoa(v)
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i]+vs[j] > vs[j]+vs[i]
	})
	if vs[0] == "0" {
		return "0"
	}
	return strings.Join(vs, "")
}

// Test cases

func main() {
	// Test cases
	nums := []int{10, 2}
	println(largestNumber(nums)) // "210"
	nums = []int{3, 30, 34, 5, 9}
	println(largestNumber(nums)) // "9534330"
}

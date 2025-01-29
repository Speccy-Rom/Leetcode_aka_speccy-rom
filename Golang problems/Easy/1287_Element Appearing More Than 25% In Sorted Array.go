/*
Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.



Example 1:

Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6
Example 2:

Input: arr = [1,1]
Output: 1


Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 105
*/

package main

func findSpecialInteger(arr []int) int {
	n := len(arr)
	for i, val := range arr {
		if val == arr[i+(n>>2)] {
			return val
		}
	}
	return 0
}

func main() {
	println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	println(findSpecialInteger([]int{1, 1}))
}

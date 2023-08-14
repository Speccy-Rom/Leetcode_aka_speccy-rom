/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
package main

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSort(nums, 0, n-1, n-k)
}

func quickSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	i, j := left-1, right+1
	x := nums[(left+right)>>1]
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if j < k {
		return quickSort(nums, j+1, right, k)
	}
	return quickSort(nums, left, j, k)
}

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)

}

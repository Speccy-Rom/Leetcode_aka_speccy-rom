/*
You are given an array nums of n positive integers.

You can perform two types of operations on any element of the array any number of times:

If the element is even, divide it by 2.
For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
If the element is odd, multiply it by 2.
For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
The deviation of the array is the maximum difference between any two elements in the array.

Return the minimum deviation the array can have after performing some number of operations.



Example 1:

Input: nums = [1,2,3,4]
Output: 1
Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
Example 2:

Input: nums = [4,1,5,20,3]
Output: 3
Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
Example 3:

Input: nums = [2,10,8]
Output: 3


Constraints:

n == nums.length
2 <= n <= 5 * 104
1 <= nums[i] <= 109
*/

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() interface{} {
	if len(*h) == 0 {
		return nil
	}
	return (*h)[0]
}

func minimumDeviation(nums []int) int {
	pq := &IntHeap{}
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		heap.Push(pq, -num)
	}

	ans := int(1e9)
	mi := -pq.Peek().(int)
	for {
		num := -heap.Pop(pq).(int)
		ans = min(ans, num-mi)
		if num%2 == 1 {
			break
		}
		mi = min(mi, num/2)
		heap.Push(pq, -num/2)
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
	nums := []int{1, 2, 3, 4}
	fmt.Println(minimumDeviation(nums)) // Output: 1

	nums = []int{4, 1, 5, 20, 3}
	fmt.Println(minimumDeviation(nums)) // Output: 3

	nums = []int{2, 10, 8}
	fmt.Println(minimumDeviation(nums)) // Output: 3
}

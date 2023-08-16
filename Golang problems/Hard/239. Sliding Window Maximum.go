/*

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length

*/

package main

import "container/heap"

func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := hp{}
	for i, v := range nums[:k-1] {
		heap.Push(&q, pair{v, i})
	}
	for i := k - 1; i < len(nums); i++ {
		heap.Push(&q, pair{nums[i], i})
		for q[0].i <= i-k {
			heap.Pop(&q)
		}
		ans = append(ans, q[0].v)
	}
	return
}

type pair struct{ v, i int }

type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.v > b.v || (a.v == b.v && i < j)
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	maxSlidingWindow([]int{1}, 1)

}

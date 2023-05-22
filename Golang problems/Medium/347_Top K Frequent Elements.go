/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	h := hp{}
	for v, freq := range cnt {
		heap.Push(&h, pair{v, freq})
		if len(h) > k {
			heap.Pop(&h)
		}
	}
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).v
	}
	return ans
}

type pair struct{ v, cnt int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].cnt < h[j].cnt }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

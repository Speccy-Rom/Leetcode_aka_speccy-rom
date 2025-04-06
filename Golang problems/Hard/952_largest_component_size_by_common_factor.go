/*
You are given an integer array of unique positive integers nums. Consider the following graph:

There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share a common factor greater than 1.
Return the size of the largest connected component in the graph.

 https://i.imgur.com/UoSKH42.jpeg

Example 1:


Input: nums = [4,6,15,35]
Output: 4
Example 2:


Input: nums = [20,50,9,63]
Output: 2
Example 3:


Input: nums = [2,3,6,7,4,12,21,39]
Output: 8


Constraints:

1 <= nums.length <= 2 * 104
1 <= nums[i] <= 105
All the values of nums are unique.
*/

package main

import (
	"fmt"
)

// UnionFind represents a disjoint set data structure
type UnionFind struct {
	parent []int
	rank   []int
	size   []int
}

// NewUnionFind creates a new UnionFind structure
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Find returns the root of the set containing x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	uf.size[px] += uf.size[py]
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
}

// GetSize returns the size of the set containing x
func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

// getPrimeFactors returns all prime factors of a number
func getPrimeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// largestComponentSize finds the size of the largest connected component
// in the graph where nodes are connected if they share a common factor > 1
func largestComponentSize(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a map to store the first occurrence of each prime factor
	factorToIndex := make(map[int]int)
	uf := NewUnionFind(len(nums))

	// For each number, connect it to all its prime factors
	for i, num := range nums {
		factors := getPrimeFactors(num)
		for _, factor := range factors {
			if prevIndex, exists := factorToIndex[factor]; exists {
				uf.Union(i, prevIndex)
			} else {
				factorToIndex[factor] = i
			}
		}
	}

	// Find the largest component size
	maxSize := 0
	for i := range nums {
		size := uf.GetSize(i)
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}

func main() {
	// Test cases
	testCases := [][]int{
		{4, 6, 15, 35},
		{20, 50, 9, 63},
		{2, 3, 6, 7, 4, 12, 21, 39},
		{1},
		{2, 3, 5, 7, 11, 13, 17, 19},
		{4, 8, 16, 32, 64},
	}

	expectedResults := []int{
		4, // [4,6,15,35] - все числа связаны через общие множители
		2, // [20,50,9,63] - только 20 и 50 связаны
		8, // [2,3,6,7,4,12,21,39] - все числа связаны
		1, // [1] - одно число
		1, // [2,3,5,7,11,13,17,19] - простые числа не связаны
		5, // [4,8,16,32,64] - все числа связаны через степени двойки
	}

	for i, nums := range testCases {
		result := largestComponentSize(nums)
		expected := expectedResults[i]
		fmt.Printf("Test case %d:\n", i+1)
		fmt.Printf("Input: %v\n", nums)
		fmt.Printf("Output: %d\n", result)
		fmt.Printf("Expected: %d\n\n", expected)
	}
}

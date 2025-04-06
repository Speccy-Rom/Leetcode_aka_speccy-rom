package medium

import (
	"fmt"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3},
		{1, 2, 4, 8},
		{1},
		{2, 3, 4, 9, 8},
		{3, 4, 16, 8},
	}

	for i, nums := range testCases {
		result := largestDivisibleSubset(nums)
		fmt.Printf("Test case %d:\n", i+1)
		fmt.Printf("Input: %v\n", nums)
		fmt.Printf("Output: %v\n\n", result)
	}
}

package hard

import (
	"fmt"
	"testing"
)

func TestLargestComponentSize(t *testing.T) {
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

		if result != expected {
			t.Errorf("Test case %d failed: got %d, expected %d", i+1, result, expected)
		}
	}
}

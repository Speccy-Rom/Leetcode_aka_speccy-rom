/*
You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.

Your solution must run in O(log n) time and O(1) space.



Example 1:

Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:

Input: nums = [3,3,7,7,10,11,11]
Output: 10


Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 105

*/

/*
Данная задача можно решить с помощью алгоритма бинарного поиска (binary search) за O(log n) времени и O(1) памяти.

Идея алгоритма состоит в том, чтобы сравнивать средний элемент массива с его соседями, чтобы определить,
где находится искомый элемент.
Если средний элемент является уникальным, мы можем вернуть его.
Если средний элемент является парным, мы можем проверить, находится ли уникальный элемент слева
или справа от него, и затем продолжить поиск только в этой половине массива.
*/

package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if mid%2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			right = mid
		} else {
			left = mid + 2
		}
	}

	return nums[left]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Print(singleNonDuplicate(nums))
}

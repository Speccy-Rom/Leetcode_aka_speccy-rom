/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.



Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23


Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109

*/

package main

func minEatingSpeed(piles []int, h int) int {
	// find the max value in piles
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	// binary search
	left, right := 1, max
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFinish(piles []int, h, k int) bool {
	time := 0
	for _, p := range piles {
		time += p / k
		if p%k != 0 {
			time++
		}
	}
	return time <= h
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	minEatingSpeed(piles, h)
}

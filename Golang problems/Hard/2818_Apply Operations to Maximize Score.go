/*
You are given an array nums of n positive integers and an integer k.

Initially, you start with a score of 1. You have to maximize your score by applying the following operation at most k times:

Choose any non-empty subarray nums[l, ..., r] that you haven't chosen previously.
Choose an element x of nums[l, ..., r] with the highest prime score. If multiple such elements exist, choose the one with the smallest index.
Multiply your score by x.
Here, nums[l, ..., r] denotes the subarray of nums starting at index l and ending at the index r, both ends being inclusive.

The prime score of an integer x is equal to the number of distinct prime factors of x. For example, the prime score of 300 is 3 since 300 = 2 * 2 * 3 * 5 * 5.

Return the maximum possible score after applying at most k operations.

Since the answer may be large, return it modulo 109 + 7.



Example 1:

Input: nums = [8,3,9,3,8], k = 2
Output: 81
Explanation: To get a score of 81, we can apply the following operations:
- Choose subarray nums[2, ..., 2]. nums[2] is the only element in this subarray. Hence, we multiply the score by nums[2]. The score becomes 1 * 9 = 9.
- Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 1, but nums[2] has the smaller index. Hence, we multiply the score by nums[2]. The score becomes 9 * 9 = 81.
It can be proven that 81 is the highest score one can obtain.
Example 2:

Input: nums = [19,12,14,6,10,18], k = 3
Output: 4788
Explanation: To get a score of 4788, we can apply the following operations:
- Choose subarray nums[0, ..., 0]. nums[0] is the only element in this subarray. Hence, we multiply the score by nums[0]. The score becomes 1 * 19 = 19.
- Choose subarray nums[5, ..., 5]. nums[5] is the only element in this subarray. Hence, we multiply the score by nums[5]. The score becomes 19 * 18 = 342.
- Choose subarray nums[2, ..., 3]. Both nums[2] and nums[3] have a prime score of 2, but nums[2] has the smaller index. Hence, we multipy the score by nums[2]. The score becomes 342 * 14 = 4788.
It can be proven that 4788 is the highest score one can obtain.


Constraints:

1 <= nums.length == n <= 105
1 <= nums[i] <= 105
1 <= k <= min(n * (n + 1) / 2, 109)

*/

package main

import "sort"

func maximumScore(nums []int, k int) int {
	n := len(nums)
	const mod = 1e9 + 7
	qpow := func(a, n int) int {
		ans := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ans = ans * a % mod
			}
			a = a * a % mod
		}
		return ans
	}
	arr := make([][3]int, n)
	left := make([]int, n)
	right := make([]int, n)
	for i, x := range nums {
		left[i] = -1
		right[i] = n
		arr[i] = [3]int{i, primeFactors(x), x}
	}
	stk := []int{}
	for _, e := range arr {
		i, f := e[0], e[1]
		for len(stk) > 0 && arr[stk[len(stk)-1]][1] < f {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		f := arr[i][1]
		for len(stk) > 0 && arr[stk[len(stk)-1]][1] <= f {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][2] > arr[j][2] })
	ans := 1
	for _, e := range arr {
		i, x := e[0], e[2]
		l, r := left[i], right[i]
		cnt := (i - l) * (r - i)
		if cnt <= k {
			ans = ans * qpow(x, cnt) % mod
			k -= cnt
		} else {
			ans = ans * qpow(x, k) % mod
			break
		}
	}
	return ans
}

func primeFactors(n int) int {
	i := 2
	ans := map[int]bool{}
	for i <= n/i {
		for n%i == 0 {
			ans[i] = true
			n /= i
		}
		i++
	}
	if n > 1 {
		ans[n] = true
	}
	return len(ans)
}

func main() {
	// Example usage
	nums := []int{8, 3, 9, 3, 8}
	k := 2
	result := maximumScore(nums, k)
	println(result) // Output: 81

	nums = []int{19, 12, 14, 6, 10, 18}
	k = 3
	result = maximumScore(nums, k)
	println(result) // Output: 4788
}

/*
There is a bookstore owner that has a store open for n minutes. Every minute, some number of customers enter the store. You are given an integer array customers of length n where customers[i] is the number of the customer that enters the store at the start of the ith minute and all those customers leave after the end of that minute.

On some minutes, the bookstore owner is grumpy. You are given a binary array grumpy where grumpy[i] is 1 if the bookstore owner is grumpy during the ith minute, and is 0 otherwise.

When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise, they are satisfied.

The bookstore owner knows a secret technique to keep themselves not grumpy for minutes consecutive minutes, but can only use it once.

Return the maximum number of customers that can be satisfied throughout the day.



Example 1:

Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes = 3
Output: 16
Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes.
The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.
Example 2:

Input: customers = [1], grumpy = [0], minutes = 1
Output: 1


Constraints:

n == customers.length == grumpy.length
1 <= minutes <= n <= 2 * 104
0 <= customers[i] <= 1000
grumpy[i] is either 0 or 1.
*/

package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var cnt, tot int
	for i, c := range customers[:minutes] {
		cnt += c * grumpy[i]
		tot += c * (grumpy[i] ^ 1)
	}
	mx := cnt
	for i := minutes; i < len(customers); i++ {
		cnt += customers[i] * grumpy[i]
		cnt -= customers[i-minutes] * grumpy[i-minutes]
		mx = max(mx, cnt)
		tot += customers[i] * (grumpy[i] ^ 1)
	}
	return tot + mx
}

// Test cases
func main() {
	// Test cases
	res := maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)
	fmt.Println(res)
	res = maxSatisfied([]int{1}, []int{0}, 1)
	fmt.Println(res)
}

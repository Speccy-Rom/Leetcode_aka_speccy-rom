/*
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.



Example 1:

Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
Example 2:

Input: salary = [1000,2000,3000]
Output: 2000.00000
Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
Average salary excluding minimum and maximum salary is (2000) / 1 = 2000


Constraints:

3 <= salary.length <= 100
1000 <= salary[i] <= 106
All the integers of salary are unique.
*/

package main

import "fmt"

func average(salary []int) float64 {
	s := 0
	mi, mx := 10000000, 0
	for _, v := range salary {
		s += v
		mi = min1491(mi, v)
		mx = max1491(mx, v)
	}
	s -= (mi + mx)
	return float64(s) / float64(len(salary)-2)
}

func max1491(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min1491(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 3000}))
}

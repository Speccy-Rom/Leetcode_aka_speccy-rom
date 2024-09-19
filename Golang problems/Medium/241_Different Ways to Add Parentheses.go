/*
Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.

The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 104.



Example 1:

Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
Example 2:

Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

*/

package main

import "strconv"

var memo = map[string][]int{}

func diffWaysToCompute(expression string) []int {
	return dfs241(expression)
}

func dfs241(exp string) []int {
	if v, ok := memo[exp]; ok {
		return v
	}
	if len(exp) < 3 {
		v, _ := strconv.Atoi(exp)
		return []int{v}
	}
	var ans []int
	for i, c := range exp {
		if c == '-' || c == '+' || c == '*' {
			left, right := dfs241(exp[:i]), dfs241(exp[i+1:])
			for _, a := range left {
				for _, b := range right {
					if c == '-' {
						ans = append(ans, a-b)
					} else if c == '+' {
						ans = append(ans, a+b)
					} else {
						ans = append(ans, a*b)
					}
				}
			}
		}
	}
	memo[exp] = ans
	return ans
}

// Test cases

func main() {
	// Test cases
	println(diffWaysToCompute("2-1-1"))   // [0,2]
	println(diffWaysToCompute("2*3-4*5")) // [-34,-14,-10,-10,10]

}

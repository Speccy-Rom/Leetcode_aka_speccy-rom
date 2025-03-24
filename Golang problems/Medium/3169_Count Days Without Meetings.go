/*
You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1). You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

Return the count of days when the employee is available for work but no meetings are scheduled.

Note: The meetings may overlap.



Example 1:

Input: days = 10, meetings = [[5,7],[1,3],[9,10]]

Output: 2

Explanation:

There is no meeting scheduled on the 4th and 8th days.

Example 2:

Input: days = 5, meetings = [[2,4],[1,3]]

Output: 1

Explanation:

There is no meeting scheduled on the 5th day.

Example 3:

Input: days = 6, meetings = [[1,6]]

Output: 0

Explanation:

Meetings are scheduled for all working days.



Constraints:

1 <= days <= 109
1 <= meetings.length <= 105
meetings[i].length == 2
1 <= meetings[i][0] <= meetings[i][1] <= days

*/

package main

import "sort"

func countDays(days int, meetings [][]int) (ans int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	last := 0
	for _, e := range meetings {
		st, ed := e[0], e[1]
		if last < st {
			ans += st - last - 1
		}
		last = max(last, ed)
	}
	ans += days - last
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage
	days := 10
	meetings := [][]int{{5, 7}, {1, 3}, {9, 10}}
	result := countDays(days, meetings)
	println(result) // Output: 2

	days = 5
	meetings = [][]int{{2, 4}, {1, 3}}
	result = countDays(days, meetings)
	println(result) // Output: 1

	days = 6
	meetings = [][]int{{1, 6}}
	result = countDays(days, meetings)
	println(result) // Output: 0

}

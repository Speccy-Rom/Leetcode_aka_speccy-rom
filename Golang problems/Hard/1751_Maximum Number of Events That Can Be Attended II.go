/*
You are given an array of events where events[i] = [startDayi, endDayi, valuei]. The ith event starts at startDayi and ends at endDayi, and if you attend this event, you will receive a value of valuei. You are also given an integer k which represents the maximum number of events you can attend.

You can only attend one event at a time. If you choose to attend an event, you must attend the entire event. Note that the end day is inclusive: that is, you cannot attend two events where one of them starts and the other ends on the same day.

Return the maximum sum of values that you can receive by attending events.



Example 1:



Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
Output: 7
Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.
Example 2:



Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
Output: 10
Explanation: Choose event 2 for a total value of 10.
Notice that you cannot attend any other event as they overlap, and that you do not have to attend k events.
Example 3:



Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
Output: 9
Explanation: Although the events do not overlap, you can only attend 3 events. Pick the highest valued three.


Constraints:

1 <= k <= events.length
1 <= k * events.length <= 106
1 <= startDayi <= endDayi <= 109
1 <= valuei <= 106
*/

package main

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	n := len(events)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	var dfs func(i, k int) int
	dfs = func(i, k int) int {
		if i >= n || k <= 0 {
			return 0
		}
		if f[i][k] > 0 {
			return f[i][k]
		}
		j := sort.Search(n, func(h int) bool { return events[h][0] > events[i][1] })
		ans := max(dfs(i+1, k), dfs(j, k-1)+events[i][2])
		f[i][k] = ans
		return ans
	}
	return dfs(0, k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2))
	fmt.Println(maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}, 2))
}

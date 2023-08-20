/*

There are n items each belonging to zero or one of m groups where group[i] is the group that the i-th item belongs to and it's equal to -1 if the i-th item belongs to no group. The items and the groups are zero indexed. A group can have no item belonging to it.

Return a sorted list of the items such that:

The items that belong to the same group are next to each other in the sorted list.
There are some relations between these items where beforeItems[i] is a list containing all the items that should come before the i-th item in the sorted array (to the left of the i-th item).
Return any solution if there is more than one solution and return an empty list if there is no solution.



Example 1:



Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
Output: [6,3,4,1,5,2,0,7]
Example 2:

Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
Output: []
Explanation: This is the same as example 1 except that 4 needs to be before 6 in the sorted list.


Constraints:

1 <= m <= n <= 3 * 104
group.length == beforeItems.length == n
-1 <= group[i] <= m - 1
0 <= beforeItems[i].length <= n - 1
0 <= beforeItems[i][j] <= n - 1
i != beforeItems[i][j]
beforeItems[i] does not contain duplicates elements.



*/

package main

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	idx := m
	groupItems := make([][]int, n+m)
	itemDegree := make([]int, n)
	groupDegree := make([]int, n+m)
	itemGraph := make([][]int, n)
	groupGraph := make([][]int, n+m)
	for i, g := range group {
		if g == -1 {
			group[i] = idx
			idx++
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}
	for i, gi := range group {
		for _, j := range beforeItems[i] {
			gj := group[j]
			if gi == gj {
				itemDegree[i]++
				itemGraph[j] = append(itemGraph[j], i)
			} else {
				groupDegree[gi]++
				groupGraph[gj] = append(groupGraph[gj], gi)
			}
		}
	}
	items := make([]int, n+m)
	for i := range items {
		items[i] = i
	}
	topoSort := func(degree []int, graph [][]int, items []int) []int {
		q := []int{}
		for _, i := range items {
			if degree[i] == 0 {
				q = append(q, i)
			}
		}
		ans := []int{}
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			ans = append(ans, i)
			for _, j := range graph[i] {
				degree[j]--
				if degree[j] == 0 {
					q = append(q, j)
				}
			}
		}
		return ans
	}
	groupOrder := topoSort(groupDegree, groupGraph, items)
	if len(groupOrder) != len(items) {
		return nil
	}
	ans := []int{}
	for _, gi := range groupOrder {
		items = groupItems[gi]
		itemOrder := topoSort(itemDegree, itemGraph, items)
		if len(items) != len(itemOrder) {
			return nil
		}
		ans = append(ans, itemOrder...)
	}
	return ans
}

func main() {
	sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}})
	sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}})

}

/*
There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.

You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.

A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge from xi to xi+1 for every 1 <= i < k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.

Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.



Example 1:



Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
Output: 3
Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
Example 2:



Input: colors = "a", edges = [[0,0]]
Output: -1
Explanation: There is a cycle from 0 to 0.


Constraints:

n == colors.length
m == edges.length
1 <= n <= 105
0 <= m <= 105
colors consists of lowercase English letters.
0 <= aj, bj < n
*/

package main

import "fmt"

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	graph := make([][]int, n)
	inDegree := make([]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		inDegree[edge[1]]++
	}
	// Topological sort
	queue := make([]int, 0)
	for i, _ := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// BFS
	// colorCount[i][j] = number of nodes with color j in the path from node i to the end
	colorCount := make([][]int, n)
	for i, _ := range colorCount {
		colorCount[i] = make([]int, 26)
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		colorCount[u][colors[u]-'a']++
		for _, v := range graph[u] {
			for i, _ := range colorCount[v] {
				colorCount[v][i] = max(colorCount[v][i], colorCount[u][i])
			}
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}

	}

	// Check if there is a cycle
	for _, d := range inDegree {
		if d > 0 {
			return -1
		}
	}

	// Find the largest color value
	ans := 0
	for _, c := range colorCount {
		for _, v := range c {
			ans = max(ans, v)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
	fmt.Println(largestPathValue("a", [][]int{{0, 0}}))

}

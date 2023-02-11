/*

You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1. Each edge is red or blue in this graph, and there could be self-edges and parallel edges.

You are given two arrays redEdges and blueEdges where:

redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi in the graph, and
blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node vj in the graph.
Return an array answer of length n, where each answer[x] is the length of the shortest path from node 0 to node x such that the edge colors alternate along the path, or -1 if such a path does not exist.



Example 1:

Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
Output: [0,1,-1]
Example 2:

Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
Output: [0,1,-1]


Constraints:

1 <= n <= 100
0 <= redEdges.length, blueEdges.length <= 400
redEdges[i].length == blueEdges[j].length == 2
0 <= ai, bi, uj, vj < n

*/

package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	red := make([][]int, n)
	blue := make([][]int, n)
	for _, e := range redEdges {
		red[e[0]] = append(red[e[0]], e[1])
	}
	for _, e := range blueEdges {
		blue[e[0]] = append(blue[e[0]], e[1])
	}
	redDist := make([]int, n)
	blueDist := make([]int, n)
	for i := 0; i < n; i++ {
		redDist[i] = -1
		blueDist[i] = -1
	}
	redDist[0] = 0
	blueDist[0] = 0
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	queue = append(queue, []int{0, 1})
	for len(queue) > 0 {
		node := queue[0][0]
		color := queue[0][1]
		queue = queue[1:]
		if color == 0 {
			for _, e := range red[node] {
				if redDist[e] == -1 {
					redDist[e] = redDist[node] + 1
					queue = append(queue, []int{e, 1})
				}
			}
		} else {
			for _, e := range blue[node] {
				if blueDist[e] == -1 {
					blueDist[e] = blueDist[node] + 1
					queue = append(queue, []int{e, 0})
				}
			}
		}
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if redDist[i] == -1 {
			result[i] = blueDist[i]
		} else if blueDist[i] == -1 {
			result[i] = redDist[i]
		} else {
			result[i] = min(redDist[i], blueDist[i])
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func main() {

}

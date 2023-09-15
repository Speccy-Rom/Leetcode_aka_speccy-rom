/*

You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].

The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value of val.

Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.



Example 1:


Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation:

We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.
Example 2:

Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18


Constraints:

1 <= points.length <= 1000
-106 <= xi, yi <= 106
All pairs (xi, yi) are distinct.

*/

package main

func minCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	g := make([][]int, n)
	vis := make([]bool, n)
	dist := make([]int, n)
	for i := range g {
		g[i] = make([]int, n)
		dist[i] = 1 << 30
	}
	for i := range g {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			t := abs(x1-x2) + abs(y1-y2)
			g[i][j] = t
			g[j][i] = t
		}
	}
	dist[0] = 0
	for i := 0; i < n; i++ {
		j := -1
		for k := 0; k < n; k++ {
			if !vis[k] && (j == -1 || dist[k] < dist[j]) {
				j = k
			}
		}
		vis[j] = true
		ans += dist[j]
		for k := 0; k < n; k++ {
			if !vis[k] {
				dist[k] = min(dist[k], g[j][k])
			}
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})
	minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}})
	minCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}})

}

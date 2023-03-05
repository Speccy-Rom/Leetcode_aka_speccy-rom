/*
Given an n x n grid containing only values 0 and 1, where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized, and return the distance. If no land or water exists in the grid, return -1.

The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.



Example 1:


Input: grid = [[1,0,1],[0,0,0],[1,0,1]]
Output: 2
Explanation: The cell (1, 1) is as far as possible from all the land with distance 2.
Example 2:


Input: grid = [[1,0,0],[0,0,0],[0,0,0]]
Output: 4
Explanation: The cell (2, 2) is as far as possible from all the land with distance 4.


Constraints:

n == grid.length
n == grid[i].length
1 <= n <= 100
grid[i][j] is 0 or 1
*/

package main

func maxDistance(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	n := len(grid)
	q := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	if len(q) == 0 || len(q) == n*n {
		return -1
	}
	res := -1
	for len(q) > 0 {
		res++
		for size := len(q); size > 0; size-- {
			cur := q[0]
			q = q[1:]
			for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				x := cur[0] + d[0]
				y := cur[1] + d[1]
				if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 0 {
					continue
				}
				grid[x][y] = 1
				q = append(q, []int{x, y})
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	println(maxDistance(grid))

}

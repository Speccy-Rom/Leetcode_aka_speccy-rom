/*

You are given an m x n matrix of characters box representing a side-view of a box. Each cell of the box is one of the following:

A stone '#'
A stationary obstacle '*'
Empty '.'
The box is rotated 90 degrees clockwise, causing some of the stones to fall due to gravity. Each stone falls down until it lands on an obstacle, another stone, or the bottom of the box. Gravity does not affect the obstacles' positions, and the inertia from the box's rotation does not affect the stones' horizontal positions.

It is guaranteed that each stone in box rests on an obstacle, another stone, or the bottom of the box.

Return an n x m matrix representing the box after the rotation described above.

*/

package main

import (
	"fmt"
	"reflect"
)

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[j][m-i-1] = box[i][j]
		}
	}
	for j := 0; j < m; j++ {
		q := []int{}
		for i := n - 1; i >= 0; i-- {
			if ans[i][j] == '*' {
				q = []int{}
			} else if ans[i][j] == '.' {
				q = append(q, i)
			} else if len(q) > 0 {
				ans[q[0]][j] = '#'
				q = q[1:]
				ans[i][j] = '.'
				q = append(q, i)
			}
		}
	}
	return ans
}

func main() {
	// Test cases
	tests := []struct {
		box  [][]byte
		want [][]byte
	}{
		{[][]byte{{'#', '.', '#'}}, [][]byte{{'.'}, {'#'}, {'#'}}},
		{[][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}, [][]byte{{'.', '#'}, {'.', '#'}, {'#', '*'}, {'#', '.'}}},
		{[][]byte{{'#', '#', '*', '.', '*', '.'}, {'#', '#', '#', '*', '.', '.'}, {'#', '#', '#', '.', '#', '.'}}, [][]byte{{'.', '#', '#'}, {'.', '#', '#'}, {'#', '#', '*'}, {'#', '*', '.'}, {'#', '.', '*'}, {'#', '.', '.'}}},
	}
	for _, t := range tests {
		if got := rotateTheBox(t.box); !reflect.DeepEqual(got, t.want) {
			fmt.Printf("rotateTheBox(%v) = %v; want %v\n", t.box, got, t.want)
		}
	}
}

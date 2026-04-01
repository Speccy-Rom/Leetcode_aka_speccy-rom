package main

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return positions[idx[i]] < positions[idx[j]]
	})

	stk := []int{}

	for _, i := range idx {
		if directions[i] == 'R' {
			stk = append(stk, i)
			continue
		}

		for len(stk) > 0 && healths[i] > 0 {
			j := stk[len(stk)-1]

			if healths[j] > healths[i] {
				healths[j]--
				healths[i] = 0
			} else if healths[j] < healths[i] {
				healths[i]--
				healths[j] = 0
				stk = stk[:len(stk)-1]
			} else {
				healths[i], healths[j] = 0, 0
				stk = stk[:len(stk)-1]
				break
			}
		}
	}

	ans := []int{}
	for _, h := range healths {
		if h > 0 {
			ans = append(ans, h)
		}
	}
	return ans
}

func main() {
	survivedRobotsHealths([]int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR")
}

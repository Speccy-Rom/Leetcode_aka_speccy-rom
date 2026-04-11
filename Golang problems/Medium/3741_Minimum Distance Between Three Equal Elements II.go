package main

func minimumDistance(nums []int) int {
	g := make(map[int][]int)
	for i, x := range nums {
		g[x] = append(g[x], i)
	}

	inf := 1 << 30
	ans := inf

	for _, ls := range g {
		m := len(ls)
		for h := 0; h < m-2; h++ {
			i := ls[h]
			k := ls[h+2]
			t := (k - i) * 2
			ans = minInt(ans, t)
		}
	}

	if ans == inf {
		return -1
	}
	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minimumDistance([]int{1, 2, 1, 1, 3})
}

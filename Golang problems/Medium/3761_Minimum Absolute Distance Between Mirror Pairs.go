package main

func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	pos := map[int]int{}
	ans := n + 1

	reverse := func(x int) int {
		y := 0
		for ; x > 0; x /= 10 {
			y = y*10 + x%10
		}
		return y
	}

	for i, x := range nums {
		if j, ok := pos[x]; ok {
			ans = minInt(ans, i-j)
		}
		pos[reverse(x)] = i
	}

	if ans > n {
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
	minMirrorPairDistance([]int{12, 21, 34, 43})
}

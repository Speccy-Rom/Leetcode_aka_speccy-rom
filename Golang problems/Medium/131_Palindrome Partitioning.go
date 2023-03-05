package main

import "fmt"

//# Given a string s, partition s such that every
//substring
// of the partition is a
//palindrome
//. Return all possible palindrome partitioning of s.
//
//
//
//Example 1:
//
//Input: s = "aab"
//Output: [["a","a","b"],["aa","b"]]
//Example 2:
//
//Input: s = "a"
//Output: [["a"]]

//func partition(s string) [][]string {
//	var result [][]string
//	var path []string
//	backtrack(s, 0, path, &result)
//	return result
//
//}
//
//func backtrack(s string, start int, path []string, result *[][]string) {
//	if start == len(s) {
//		*result = append(*result, append([]string{}, path...))
//		return
//	}
//	for i := start; i < len(s); i++ {
//		if isPalindrome(s, start, i) {
//			path = append(path, s[start:i+1])
//			backtrack(s, i+1, path, result)
//			path = path[:len(path)-1]
//		}
//	}
//}
//
//func isPalindrome(s string, start, end int) bool {
//	for start < end {
//		if s[start] != s[end] {
//			return false
//		}
//		start++
//		end--
//	}
//	return true
//}

func partition(s string) [][]string {
	var ans [][]string
	runes := []rune(s)

	palindrome := func(b, e int) bool {
		for ; b < e; b, e = b+1, e-1 {
			if runes[b] != runes[e] {
				return false
			}
		}
		return true
	}
	var dfs func(int, []string)
	dfs = func(pos int, seq []string) {
		if pos == len(runes) {
			t := make([]string, len(seq))
			copy(t, seq)
			ans = append(ans, t)
			return
		}
		for i := pos + 1; i <= len(runes); i++ {
			if palindrome(pos, i-1) {
				seq = append(seq, string(runes[pos:i]))
				dfs(i, seq)
				seq = seq[:len(seq)-1]
			}
		}
	}
	dfs(0, nil)
	return ans

}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))

}

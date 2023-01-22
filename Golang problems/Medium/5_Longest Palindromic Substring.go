package main

// Given a string s, return the longest palindromic substring in s.
//
//
//
//Example 1:
//
//Input: s = "babad"
//Output: "bab"
//Explanation: "aba" is also a valid answer.
//Example 2:
//
//Input: s = "cbbd"
//Output: "bb"
//
//
//Constraints:
//
//1 <= s.length <= 1000
//s consist of only digits and English letters.

func longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "babad"
	println(longestPalindrome(s))
}

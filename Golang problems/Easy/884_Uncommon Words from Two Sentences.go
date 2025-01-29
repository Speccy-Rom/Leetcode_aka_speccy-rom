/*
A sentence is a string of single-space separated words where each word consists only of lowercase letters.

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.



Example 1:

Input: s1 = "this apple is sweet", s2 = "this apple is sour"

Output: ["sweet","sour"]

Explanation:

The word "sweet" appears only in s1, while the word "sour" appears only in s2.

Example 2:

Input: s1 = "apple apple", s2 = "banana"

Output: ["banana"]


*/

package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) (ans []string) {
	cnt := map[string]int{}
	for _, s := range strings.Split(s1, " ") {
		cnt[s]++
	}
	for _, s := range strings.Split(s2, " ") {
		cnt[s]++
	}
	for s, v := range cnt {
		if v == 1 {
			ans = append(ans, s)
		}
	}
	return
}

func main() {
	// Test cases
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2)) // ["sweet","sour"]
	s1 = "apple apple"
	s2 = "banana"
	fmt.Println(uncommonFromSentences(s1, s2)) // ["banana"]
}

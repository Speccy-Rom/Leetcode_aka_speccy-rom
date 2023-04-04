/*Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".


Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.
*/

package main

//func findAnagrams(s string, p string) []int {
//	var res []int
//	pmap := make(map[byte]int)
//	for i := 0; i < len(p); i++ {
//		pmap[p[i]]++
//	}
//	for i := 0; i < len(s)-len(p)+1; i++ {
//		smap := make(map[byte]int)
//		for j := i; j < i+len(p); j++ {
//			smap[s[j]]++
//		}
//		if compare(pmap, smap) {
//			res = append(res, i)
//		}
//	}
//	return res
//}
//
//func compare(pmap, smap map[byte]int) bool {
//	for k, v := range pmap {
//		if smap[k] != v {
//			return false
//		}
//	}
//	return true
//}

func findAnagrams(s string, p string) []int {
	var result []int
	var sMask, pMask [26]int

	for _, ch := range p {
		pMask[int(ch)-int('a')] += 1
	}

	for i := 0; i <= len(s); i++ {
		if i >= len(p) {
			match := true
			for j := 0; j < 26; j++ {
				if sMask[j] != pMask[j] {
					match = false
					break
				}
			}
			if match {
				result = append(result, i-len(p))
			}
		}

		if i != len(s) {
			sMask[int(s[i])-int('a')] += 1
			if i >= len(p) {
				sMask[int(s[i-len(p)])-int('a')] -= 1
			}
		}
	}

	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	findAnagrams(s, p)
}

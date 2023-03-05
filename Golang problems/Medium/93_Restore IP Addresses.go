package main

import (
	"fmt"
	"strconv"
)

// A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
//
//For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
//Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.
//
//
//
//Example 1:
//
//Input: s = "25525511135"
//Output: ["255.255.11.135","255.255.111.35"]
//Example 2:
//
//Input: s = "0000"
//Output: ["0.0.0.0"]
//Example 3:
//
//Input: s = "101023"
//Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//Constraints:
//
//1 <= s.length <= 20
//s consists of digits only.

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var dfs func(path string, idx int, count int)
	dfs = func(path string, idx int, count int) {
		if count > 4 {
			return
		}
		if count == 4 && idx == len(s) {
			res = append(res, path)
			return
		}
		for i := 1; i <= 3; i++ {
			if idx+i > len(s) {
				break
			}
			if i != 1 && s[idx] == '0' {
				break
			}
			num, _ := strconv.Atoi(s[idx : idx+i])
			if num > 255 {
				break
			}
			if count == 0 {
				dfs(path+s[idx:idx+i], idx+i, count+1)
			} else {
				dfs(path+"."+s[idx:idx+i], idx+i, count+1)
			}
		}
	}
	dfs("", 0, 0)
	return res

}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

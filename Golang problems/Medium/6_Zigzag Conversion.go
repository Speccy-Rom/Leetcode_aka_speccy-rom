/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
*/

package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	ans := make([]byte, n)
	index := 0
	for i := 0; i < numRows; i++ {
		for j := i; j < n; j += 2 * (numRows - 1) {
			ans[index] = s[j]
			index++
			if i > 0 && i < numRows-1 && j+2*(numRows-i-1) < n {
				ans[index] = s[j+2*(numRows-i-1)]
				index++
			}
		}
	}
	return string(ans)
}

func main() {
	println(convert("PAYPALISHIRING", 3))
	println(convert("PAYPALISHIRING", 4))
	println(convert("A", 1))
	println(convert("AB", 1))
}

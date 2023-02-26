/*

Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.

*/

package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min1(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min1(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c

}

/*
Здесь мы используем динамическое программирование, чтобы решить эту задачу.
Мы создаем матрицу dp, где dp[i][j] представляет минимальное количество операций, необходимых для преобразования word1[:i] в word2[:j]. Мы заполняем матрицу с помощью следующих формул:

dp[i][0] = i для любого i, так как преобразование из пустой строки в word1[:i] требует i операций удаления.
dp[0][j] = j для любого j, так как преобразование из пустой строки в word2[:j] требует j операций вставки.
dp[i][j] = dp[i-1][j-1] если word1[i-1] == word2[j-1], так как текущие символы уже совпадают.
dp[i][j] = 1 + min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) в противном случае,
где мы рассматриваем три возможных операции - замену, вставку или удаление - и выбираем операцию с минимальной стоимостью.
Наконец, мы возвращаем dp[m][n], г

*/

func main() {
	word1 := "horse"
	word2 := "ros"
	println(minDistance(word1, word2))
}

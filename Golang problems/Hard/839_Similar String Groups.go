/*
Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y. Also two strings X and Y are similar if they are equal.

For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar, but "star" is not similar to "tars", "rats", or "arts".

Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.  Notice that "tars" and "arts" are in the same group even though they are not similar.  Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.

We are given a list strs of strings where every string in strs is an anagram of every other string in strs. How many groups are there?



Example 1:

Input: strs = ["tars","rats","arts","star"]
Output: 2
Example 2:

Input: strs = ["omv","ovm"]
Output: 1


Constraints:

1 <= strs.length <= 300
1 <= strs[i].length <= 300
strs[i] consists of lowercase letters only.
All words in strs have the same length and are anagrams of each other.
*/

// Для решения задачи мы будем использовать алгоритм DFS.
//Сначала мы создаем массив seen, чтобы отслеживать, была ли строка уже рассмотрена.
//Затем мы перебираем все строки в массиве и, если строка еще не была рассмотрена,
//запускаем DFS на этой строке, помечаем ее как просмотренную и увеличиваем счетчик групп.
//
//Функция DFS принимает массив строк A, индекс строки i и массив seen.
//Она помечает строку i как просмотренную и перебирает все строки в массиве.
//Если строка j еще не была просмотрена и похожа на строку i, то запускаем DFS на строке j.
//
//Функция isSimilar проверяет, похожи ли две строки друг на друга.
//Она считает количество символов, которые отличаются в двух строках,
//и если их количество больше двух, возвращает false.
//В противном случае возвращает true.

package main

import "fmt"

func numSimilarGroups(strs []string) int {
	ans := 0
	seen := make([]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if !seen[i] {
			dfs(strs, i, seen)
			ans++
		}
	}

	return ans
}

func dfs(strs []string, i int, seen []bool) {
	seen[i] = true
	for j := 0; j < len(strs); j++ {
		if !seen[j] && isSimilar(strs[i], strs[j]) {
			dfs(strs, j, seen)
		}
	}
}

func isSimilar(x, y string) bool {
	diff := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] && (diff+1) > 2 {
			return false
		}
		if x[i] != y[i] {
			diff++
		}
	}
	return true
}

func main() {
	strs := []string{"tars", "rats", "arts", "star"}
	fmt.Println(numSimilarGroups(strs))
	strs = []string{"omv", "ovm"}
	fmt.Println(numSimilarGroups(strs))
}

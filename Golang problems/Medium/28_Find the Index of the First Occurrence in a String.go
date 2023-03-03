/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.



Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.


Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.

*/

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	// Вычисляем префикс-функцию для подстроки needle.
	pi := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	// Проходимся по строке haystack и сравниваем символы с needle.
	j = 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

//Для решения этой задачи мы можем использовать алгоритм Кнута-Морриса-Пратта (KMP), который позволяет искать вхождения подстроки в строку с линейной сложностью O(n+m), где n - длина строки, а m - длина подстроки.
//
//Идея алгоритма KMP состоит в том, чтобы построить префикс-функцию для подстроки needle. Префикс-функция для каждой позиции i в needle указывает, какая максимальная длина суффикса needle[:i], который является также ее префиксом. Затем мы можем использовать эту префикс-функцию, чтобы определить, сколько символов строки haystack нужно пропустить, чтобы продолжить сравнение с needle, если текущие символы не совпадают.
//
//Алгоритм KMP можно реализовать следующим образом:
//
//Вычисляем префикс-функцию для подстроки needle.
//Проходимся по строке haystack и сравниваем символы с needle. Если символы совпадают, мы продолжаем сравнивать следующие символы. Если символы не совпадают, мы используем значение из префикс-функции, чтобы определить, сколько символов haystack нужно пропустить, чтобы продолжить сравнение с needle.
//Если мы находим все символы needle в haystack, мы возвращаем индекс, с которого начинается needle в haystack. Если мы не находим все символы needle в haystack, мы возвращаем -1.
//Временная сложность алгоритма KMP составляет O(n+m), где n - длина строки haystack, а m - длина подстроки needle.

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))

}

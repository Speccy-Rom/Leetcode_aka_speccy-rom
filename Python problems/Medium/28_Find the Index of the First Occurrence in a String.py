# Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
#
#
#
# Example 1:
#
# Input: haystack = "sadbutsad", needle = "sad"
# Output: 0
# Explanation: "sad" occurs at index 0 and 6.
# The first occurrence is at index 0, so we return 0.
# Example 2:
#
# Input: haystack = "leetcode", needle = "leeto"
# Output: -1
# Explanation: "leeto" did not occur in "leetcode", so we return -1.
#
#
# Constraints:
#
# 1 <= haystack.length, needle.length <= 104
# haystack and needle consist of only lowercase English characters.


# Для решения данной задачи можно использовать алгоритм Кнута-Морриса-Пратта (KMP), который работает за линейное
# время O(n+m), где n и m - длины строк haystack и needle соответственно.

# Алгоритм KMP основан на принципе сравнения символов в haystack и needle поочередно, и, если символы не совпадают,
# переходит на новый символ в haystack и продолжает поиск. При этом он использует знания о самой строке needle и
# строит для нее префикс-функцию, которая позволяет эффективно пропускать некоторые символы в haystack.
#
# Префикс-функция для строки needle позволяет найти наибольший префикс строки needle, который является ее суффиксом.
# Это позволяет определить, с какого символа в needle нужно начать сравнение с haystack, если символы не совпадают.
# Эта информация хранится в префикс-функции.
#
# Для построения префикс-функции можно использовать алгоритм КМП, который работает за O(m), где m - длина строки needle.


class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if not needle:
            return 0
        n, m = len(haystack), len(needle)
        if n < m:
            return -1
        # build prefix table
        i, j = 1, 0
        table = [0] * m
        while i < m:
            if needle[i] == needle[j]:
                j += 1
                table[i] = j
                i += 1
            elif j > 0:
                j = table[j - 1]
            else:
                i += 1
        # search for needle in haystack
        i, j = 0, 0
        while i < n:
            if haystack[i] == needle[j]:
                i += 1
                j += 1
                if j == m:
                    return i - m
            elif j > 0:
                j = table[j - 1]
            else:
                i += 1
        return -1


if __name__ == '__main__':
    res = Solution()
    print(res.strStr("sadbutsad", "sad"))
    print(res.strStr("leetcode", "leeto"))

# You are given an integer array nums and two integers minK and maxK.

# A fixed-bound subarray of nums is a subarray that satisfies the following conditions:
#
# The minimum value in the subarray is equal to minK.
# The maximum value in the subarray is equal to maxK.
# Return the number of fixed-bound subarrays.
#
# A subarray is a contiguous part of an array.
#
#
#
# Example 1:
#
# Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
# Output: 2
# Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
# Example 2:
#
# Input: nums = [1,1,1,1], minK = 1, maxK = 1
# Output: 10
# Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.
#
#
# Constraints:
#
# 2 <= nums.length <= 105
# 1 <= nums[i], minK, maxK <= 106

from typing import List


class Solution:
    def countSubarrays(self, nums: List[int], minK: int, maxK: int) -> int:
        ans = 0
        j = -1
        prevMinKIndex = -1
        prevMaxKIndex = -1
        for i, num in enumerate(nums):
            if num < minK or num > maxK:
                j = i
            if num == minK:
                prevMinKIndex = i
            if num == maxK:
                prevMaxKIndex = i
            # any index k in [j + 1, min(prevMinKIndex, prevMaxKIndex)] can be the
            # start of the subarray s.t. nums[k..i] satisfies the conditions
            ans += max(0, min(prevMinKIndex, prevMaxKIndex) - j)
        return ans

#Работа функции происходит по следующей логике:

# Инициализируем переменные: ans - количество фиксированных подмассивов, j - индекс последнего элемента,
# не превышающего maxK, prevMinKIndex и prevMaxKIndex - индексы последнего вхождения minK и maxK соответственно.
#
# Перебираем элементы массива nums и для каждого элемента проверяем, является ли он меньше minK или больше maxK.
# Если да, то обновляем значение j, указывающее на последний элемент, не превышающий maxK.
# Если нет, то переходим к следующему шагу.
# Проверяем, является ли текущий элемент равным minK или maxK. Если да, то обновляем соответствующий
# индекс последнего вхождения.
#
# Считаем количество фиксированных подмассивов, которые могут начинаться с любого
# индекса k в диапазоне [j + 1, min(prevMinKIndex, prevMaxKIndex)], где k - индекс начала подмассива, i - индекс
# текущего элемента. Это количество равно разнице между индексами prevMinKIndex и prevMaxKIndex (если оба индекса не
# равны -1) и j + 1, т.е. количеству индексов в диапазоне [j + 1, min(prevMinKIndex, prevMaxKIndex)]. Добавляем
# найденное количество подмассивов к общему числу ans. Возвращаем значение ans. Таким образом, мы проходим по всем
# элементам массива только один раз и находим ответ за линейное время.


if __name__ == '__main__':
    print(Solution().countSubarrays([1, 3, 5, 2, 7, 5], 1, 5))
    print(Solution().countSubarrays([1, 1, 1, 1], 1, 1))

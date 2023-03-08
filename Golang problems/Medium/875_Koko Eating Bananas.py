# Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone
# and will come back in h hours.

# Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k
# bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more
# bananas during this hour.
#
# Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
#
# Return the minimum integer k such that she can eat all the bananas within h hours.
#
#
#
# Example 1:
#
# Input: piles = [3,6,7,11], h = 8
# Output: 4
# Example 2:
#
# Input: piles = [30,11,23,4,20], h = 5
# Output: 30
# Example 3:
#
# Input: piles = [30,11,23,4,20], h = 6
# Output: 23
#
#
# Constraints:
#
# 1 <= piles.length <= 104
# piles.length <= h <= 109
# 1 <= piles[i] <= 109


# Для решения этой задачи мы можем использовать бинарный поиск. Мы знаем, что ответ должен лежать между 1 (
# минимальное значение k) и max(piles) (максимальное значение k). Мы можем начать с бинарного поиска в этом
# диапазоне, используя среднее значение в качестве стартовой точки. Затем мы проверяем, можно ли есть все бананы с
# помощью текущего значения k, и если да, мы уменьшаем диапазон до левой половины, иначе мы увеличиваем диапазон до
# правой половины. Мы повторяем этот процесс до тех пор, пока диапазон не сократится до одного элемента.
#
# Для проверки, можно ли есть все бананы в течение h часов с данным значением k, мы можем пройти по всем пачкам
# бананов и вычислить количество часов, необходимых для их съедения, используя формулу ceil(piles[i] / k),
# где ceil - это функция округления вверх. Если общее количество часов, необходимых для съедения всех бананов,
# меньше или равно h, это означает, что мы можем есть все бананы с использованием данного значения k.
#
# Таким образом, сложность алгоритма будет O(n log m), где n - это количество пачек бананов, а m - максимальное
# количество бананов в пачке.
import math
from typing import List


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        left, right = 1, max(piles)

        while left < right:
            mid = (left + right) // 2
            hours = sum(math.ceil(pile / mid) for pile in piles)
            if hours > h:
                left = mid + 1
            else:
                right = mid

        return left


if __name__ == '__main__':
    s = Solution()
    print(s.minEatingSpeed([3, 6, 7, 11], 8))
    print(s.minEatingSpeed([30, 11, 23, 4, 20], 5))
    print(s.minEatingSpeed([30, 11, 23, 4, 20], 6))

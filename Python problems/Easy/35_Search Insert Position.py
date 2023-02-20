# Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
#
# You must write an algorithm with O(log n) runtime complexity.
#
#
#
# Example 1:
#
# Input: nums = [1,3,5,6], target = 5
# Output: 2
# Example 2:
#
# Input: nums = [1,3,5,6], target = 2
# Output: 1
# Example 3:
#
# Input: nums = [1,3,5,6], target = 7
# Output: 4
#
#
# Constraints:
#
# 1 <= nums.length <= 104
# -104 <= nums[i] <= 104
# nums contains distinct values sorted in ascending order.
# -104 <= target <= 104
from typing import List


# Данная задача может быть решена с помощью бинарного поиска (Binary Search).
# Идея состоит в том, что мы будем делить отсортированный массив на две части до тех пор, пока не найдем индекс
# элемента, равного целевому элементу или не найдем место, куда должен быть вставлен целевой элемент.

# Алгоритм бинарного поиска следующий:
#
# Установим левую границу равной 0 и правую границу равной длине массива минус один (0 и len(nums)-1).
# Пока левая граница меньше или равна правой границе, вычисляем середину массива.
# Это делается с помощью формулы mid = left + (right - left) // 2, чтобы избежать возможной
# ошибки при больших значениях left и right.
# Если nums[mid] равно целевому элементу, возвращаем mid.
# Если nums[mid] больше целевого элемента, то целевой элемент находится слева от mid.
# Таким образом, мы устанавливаем правую границу на mid-1.
# Если nums[mid] меньше целевого элемента, то целевой элемент находится справа от mid.
# Таким образом, мы устанавливаем левую границу на mid+1.
# Если элемент не найден, мы просто возвращаем левую границу, так как это место,
# куда целевой элемент должен быть вставлен.
# Сложность алгоритма бинарного поиска в лучшем, среднем и худшем случаях - O(log n), где n - это длина массива.


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:  # временная сложность O(log n)
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return left

    # def searchInsert(self, nums: List[int], target: int) -> int:  # временная сложность O(n)
    #     for i, n in enumerate(nums):
    #         if n >= target:
    #             return i
    #     return len(nums)


if __name__ == '__main__':
    s = Solution()
    print(s.searchInsert([1, 3, 5, 6], 5))
    print(s.searchInsert([1, 3, 5, 6], 2))
    print(s.searchInsert([1, 3, 5, 6], 7))
    print(s.searchInsert([1, 3, 5, 6], 0))
    print(s.searchInsert([1], 0))

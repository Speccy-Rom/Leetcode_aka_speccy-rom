# You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
#
# Return the single element that appears only once.
#
# Your solution must run in O(log n) time and O(1) space.
#
#
#
# Example 1:
#
# Input: nums = [1,1,2,3,3,4,4,8,8]
# Output: 2
# Example 2:
#
# Input: nums = [3,3,7,7,10,11,11]
# Output: 10
#
#
# Constraints:
#
# 1 <= nums.length <= 105
# 0 <= nums[i] <= 105

#Для решения данной задачи можно использовать алгоритм бинарного поиска, поскольку массив отсортирован и имеет размерность O(log n).

# Идея состоит в том, что у нас есть пары элементов в массиве, за исключением одного, который появляется один раз.
# Если мы берем два соседних элемента и сравниваем их между собой, их индексы будут четными.
# Если два соседних элемента не совпадают, значит, искомый элемент находится в левой половине массива, иначе - в правой.
# Мы можем сократить длину массива вдвое, определив, в какой половине массива находится искомый элемент.
#
# В случае, если длина массива нечетная, мы можем сделать ее четной, удалив первый элемент или последний элемент,
# чтобы пары соответствовали всем оставшимся элементам.

from typing import List


class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        left = 0
        right = len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            if mid % 2 == 1:
                mid -= 1
            if nums[mid] == nums[mid + 1]:
                left = mid + 2
            else:
                right = mid
        return nums[left]


if __name__ == '__main__':
    solution = Solution()
    print(solution.singleNonDuplicate([1, 1, 2, 3, 3, 4, 4, 8, 8]))
    print(solution.singleNonDuplicate([3, 3, 7, 7, 10, 11, 11]))

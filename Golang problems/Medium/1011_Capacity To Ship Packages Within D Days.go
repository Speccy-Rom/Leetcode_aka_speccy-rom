/*
A conveyor belt has packages that must be shipped from one port to another within days days.

The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights). We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.



Example 1:

Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5
Output: 15
Explanation: A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.
Example 2:

Input: weights = [3,2,2,4,1,4], days = 3
Output: 6
Explanation: A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
1st day: 3, 2
2nd day: 2, 4
3rd day: 1, 4
Example 3:

Input: weights = [1,2,3,1,1], days = 4
Output: 3
Explanation:
1st day: 1
2nd day: 2
3rd day: 3
4th day: 1, 1


Constraints:

1 <= days <= weights.length <= 5 * 104
1 <= weights[i] <= 500
*/

package main

import "fmt"

// Для решения данной задачи мы можем использовать алгоритм двоичного поиска,
//который имеет сложность O(n * log(sum(weights))), где n - количество элементов в массиве weights,
//а sum(weights) - сумма всех весов.
//
//Алгоритм заключается в том, чтобы выполнять двоичный поиск по диапазону возможных значений грузоподъемности корабля.
//В качестве границ можно использовать минимально возможный вес (вес самого тяжелого груза)
//и сумму всех весов (все грузы разом).
//На каждой итерации берется среднее значение веса и проверяется, можно ли уложить все грузы
//в корабль за days дней, используя текущую грузоподъемность. Если да, то поиск продолжается
//в левой половине, иначе - в правой половине. Поиск заканчивается, когда найден минимальный
//вес корабля, при котором можно доставить все грузы в заданный срок.

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	for left < right {
		mid := left + (right-left)/2
		currDays := 1
		currSum := 0
		for _, w := range weights {
			if currSum+w > mid {
				currDays++
				currSum = 0
			}
			currSum += w
		}
		if currDays > days {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//Сначала мы находим минимальный возможный вес left и максимальный возможный вес right.
//Затем выполняем двоичный поиск, на каждой итерации проверяя, можно ли уложить все грузы
//в корабль за days дней, используя текущую грузоподъемность mid.
//Если текущее количество дней currDays больше days, то мы увеличиваем грузоподъемность
//(переходим в правую половину), иначе - уменьшаем грузоподъемность (переходим в левую половину).
//Поиск заканчивается, когда left и right сойдутся, и left будет минимальной возможной
//грузоподъемностью, при которой можно доставить все грузы в заданный срок.

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	fmt.Println(shipWithinDays(weights, days))
}

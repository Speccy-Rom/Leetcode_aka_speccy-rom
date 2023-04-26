/*
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares to Venture Capital, LeetCode would like to work on some projects to increase its capital before the IPO. Since it has limited resources, it can only finish at most k distinct projects before the IPO. Help LeetCode design the best way to maximize its total capital after finishing at most k distinct projects.

You are given n projects where the ith project has a pure profit profits[i] and a minimum capital of capital[i] is needed to start it.

Initially, you have w capital. When you finish a project, you will obtain its pure profit and the profit will be added to your total capital.

Pick a list of at most k distinct projects from given projects to maximize your final capital, and return the final maximized capital.

The answer is guaranteed to fit in a 32-bit signed integer.



Example 1:

Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
Output: 4
Explanation: Since your initial capital is 0, you can only start the project indexed 0.
After finishing it you will obtain profit 1 and your capital becomes 1.
With capital 1, you can either start the project indexed 1 or the project indexed 2.
Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.
Example 2:

Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
Output: 6


Constraints:

1 <= k <= 105
0 <= w <= 109
n == profits.length
n == capital.length
1 <= n <= 105
0 <= profits[i] <= 104
0 <= capital[i] <= 109

*/

/*
Для решения этой задачи мы будем использовать алгоритм жадной стратегии, который заключается в том, чтобы выбирать
проекты, которые могут приносить максимальную прибыль и при этом требуют минимального капитала, чтобы их реализовать.

Алгоритм:

Создаем массив проектов, содержащий пары (прибыль, капитал) для каждого проекта.
Сортируем проекты по капиталу в порядке возрастания.
Создаем приоритетную очередь для проектов, используя максимальную кучу (max heap), и добавляем все проекты,
которые можно начать в данный момент (т.е. капитал не превышает имеющийся капитал w).
Выбираем k проектов с наибольшей прибылью из приоритетной очереди и выполняем каждый проект, добавляя прибыль к имеющемуся капиталу w.
Повторяем шаги 3-4, пока не выполним k проектов или пока не закончатся проекты.
В этом алгоритме сложность сортировки проектов составляет O(nlogn), добавление и извлечение из приоритетной очереди
занимает O(logn) для каждого проекта, так что общая сложность алгоритма составляет O(nlogn + klogn).
*/

package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

type Solution struct{}

func (s *Solution) minimumDeviation(nums []int) int {
	pq := &IntHeap{}
	heap.Init(pq)

	for _, num := range nums {
		if num%2 == 0 {
			heap.Push(pq, -num)
		} else {
			heap.Push(pq, -num*2)
		}
	}

	ans := int(1e9)
	mi := -pq.Max()

	for pq.Len() > 0 {
		num := -heap.Pop(pq).(int)
		ans = min(ans, num-mi)
		if num%2 == 1 {
			break
		}
		mi = min(mi, num/2)
		heap.Push(pq, -num/2)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	println(findMaximizedCapital(k, w, profits, capital))
}

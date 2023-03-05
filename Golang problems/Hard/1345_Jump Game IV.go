/*
Given an array of integers arr, you are initially positioned at the first index of the array.

In one step you can jump from index i to index:

i + 1 where: i + 1 < arr.length.
i - 1 where: i - 1 >= 0.
j where: arr[i] == arr[j] and i != j.
Return the minimum number of steps to reach the last index of the array.

Notice that you can not jump outside of the array at any time.



Example 1:

Input: arr = [100,-23,-23,404,100,23,23,23,3,404]
Output: 3
Explanation: You need three jumps from index 0 --> 4 --> 3 --> 9. Note that index 9 is the last index of the array.
Example 2:

Input: arr = [7]
Output: 0
Explanation: Start index is the last index. You do not need to jump.
Example 3:

Input: arr = [7,6,9,6,9,6,9,7]
Output: 1
Explanation: You can jump directly from index 0 to index 7 which is last index of the array.


Constraints:

1 <= arr.length <= 5 * 104
-108 <= arr[i] <= 108
*/

package main

func minJumps(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	// Создаем словарь, где ключ - значение в массиве, а значение - массив индексов этого значения.
	indices := make(map[int][]int)
	for i, num := range arr {
		indices[num] = append(indices[num], i)
	}

	// Создаем слайс, который отслеживает посещенные индексы.
	visited := make([]bool, n)
	visited[0] = true

	// Создаем очередь, где будем хранить индексы и количество шагов до этого индекса.
	queue := []int{0}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			idx := queue[i]

			if idx == n-1 {
				return steps
			}

			// Переход на следующий элемент
			if idx+1 < n && !visited[idx+1] {
				visited[idx+1] = true
				queue = append(queue, idx+1)
			}

			// Переход на предыдущий элемент
			if idx-1 >= 0 && !visited[idx-1] {
				visited[idx-1] = true
				queue = append(queue, idx-1)
			}

			// Переход на все индексы, где значение равно значению текущего индекс
			for _, nextIdx := range indices[arr[idx]] {
				if !visited[nextIdx] {
					visited[nextIdx] = true
					queue = append(queue, nextIdx)
				}

			}

			// Удаляем все индексы, где значение равно значению текущего индекс
			indices[arr[idx]] = []int{}
		}

		steps++
		queue = queue[size:]
	}

	return -1

}

func main() {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	minJumps(arr)

}

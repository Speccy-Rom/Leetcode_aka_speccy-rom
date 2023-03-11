/*
Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

Implement the Solution class:

Solution(ListNode head) Initializes the object with the head of the singly-linked list head.
int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of the list should be equally likely to be chosen.


Example 1:


Input
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 3, 2, 2, 3]

Explanation
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // return 1
solution.getRandom(); // return 3
solution.getRandom(); // return 2
solution.getRandom(); // return 2
solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.


Constraints:

The number of nodes in the linked list will be in the range [1, 104].
-104 <= Node.val <= 104
At most 104 calls will be made to getRandom.


Follow up:

What if the linked list is extremely large and its length is unknown to you?
Could you solve this efficiently without using extra space?

*/

package main

import (
	"fmt"
	"math/rand"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	curr := this.head
	result := curr.Val
	count := 1
	for curr != nil {
		if rand.Intn(count) == 0 {
			result = curr.Val
		}
		count++
		curr = curr.Next
	}
	return result
}

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next
	solution := Constructor(head)
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
	fmt.Println(solution.GetRandom())
}

/*
Для решения данной задачи можно использовать алгоритм Reservoir Sampling с линейным временем работы O(n),
где n - количество элементов в связном списке.

Алгоритм Reservoir Sampling заключается в том, что мы проходим по всем элементам списка и выбираем из
них случайный элемент с равной вероятностью. Для этого нам нужно завести переменную result, которая будет хранить
значение выбранного элемента, а также переменную count, которая будет хранить количество элементов, которые мы уже прошли.
На каждой итерации мы увеличиваем count на 1, и с вероятностью 1/count заменяем result на текущий элемент.
В конце алгоритма result будет хранить значение выбранного случайного элемента.

Для реализации данного алгоритма в Golang, нам нужно добавить поле head в структуру Solution для хранения
головы списка, а также реализовать конструктор Constructor и метод GetRandom.

Конструктор Constructor будет просто сохранять голову списка в поле head структуры Solution.

Метод GetRandom будет проходить по всем элементам списка, применять алгоритм Reservoir Sampling и возвращать выбранное значение. Для этого мы будем использовать указатель на текущий элемент списка curr и переменную count, которая будет хранить количество пройденных элементов. Начальное значение result будет равно значению головы списка. Затем мы будем итерироваться по всем элементам списка, увеличивая count на 1 на каждой итерации и с вероятностью 1/count заменяя result на значение текущего элемента. По окончании итераций мы возвращаем result.
*/

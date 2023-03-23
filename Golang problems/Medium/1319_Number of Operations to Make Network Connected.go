/*
There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. Any computer can reach any other computer directly or indirectly through the network.

You are given an initial computer network connections. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected.

Return the minimum number of times you need to do this in order to make all the computers connected. If it is not possible, return -1.



Example 1:


Input: n = 4, connections = [[0,1],[0,2],[1,2]]
Output: 1
Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
Example 2:


Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
Output: 2
Example 3:

Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
Output: -1
Explanation: There are not enough cables.


Constraints:

1 <= n <= 105
1 <= connections.length <= min(n * (n - 1) / 2, 105)
connections[i].length == 2
0 <= ai, bi < n
ai != bi
There are no repeated connections.
No two computers are connected by more than one cable.
*/

package main

import "fmt"

func makeConnected(n int, connections [][]int) int {
	// Создаем граф с помощью списка смежности
	graph := make(map[int][]int)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	// Выполняем поиск в глубину, начиная с первой вершины графа
	visited := make(map[int]bool)
	dfs(graph, visited, 0)

	// Подсчитываем количество компонент связности графа
	numComponents := 1
	for i := 1; i < n; i++ {
		if !visited[i] {
			numComponents++
			dfs(graph, visited, i)
		}
	}

	// Подсчитываем количество ребер, которые нужно удалить
	numEdges := len(connections)
	if numEdges >= n-1 {
		return numComponents - 1
	}
	return -1
}

func dfs(graph map[int][]int, visited map[int]bool, node int) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(graph, visited, neighbor)
		}
	}
}

/*
Данная задача требует решения с помощью алгоритма поиска в глубину (DFS).
Мы можем представить данную сеть компьютеров как граф, где каждый компьютер представляет вершину, а соединения между компьютерами - это ребра графа.
Если мы можем достичь каждого компьютера от каждого другого компьютера, то граф называется связным.

Таким образом, задача сводится к поиску минимального количества ребер, которые нужно удалить из графа, чтобы сделать его связным.
Если граф уже связный, то мы не должны удалять никаких ребер.

Алгоритм:
Создать граф с помощью списка смежности на основе входных данных.
Выполнить поиск в глубину, начиная с первой вершины графа. Если мы можем достичь всех вершин графа из первой вершины, то граф уже связный, и мы не должны удалять никаких ребер.
Если граф не связный, то мы должны найти все компоненты связности графа и подсчитать количество ребер, которые нужно удалить, чтобы сделать каждую компоненту связной.
Возвращаем общее количество ребер, которые нужно удалить, чтобы сделать граф связным.
*/

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))

}

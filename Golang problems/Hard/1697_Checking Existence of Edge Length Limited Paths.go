/*
An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi] denotes an edge between nodes ui and vi with distance disi. Note that there may be multiple edges between two nodes.

Given an array queries, where queries[j] = [pj, qj, limitj], your task is to determine for each queries[j] whether there is a path between pj and qj such that each edge on the path has a distance strictly less than limitj .

Return a boolean array answer, where answer.length == queries.length and the jth value of answer is true if there is a path for queries[j] is true, and false otherwise.



Example 1:


Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
Output: [false,true]
Explanation: The above figure shows the given graph. Note that there are two overlapping edges between 0 and 1 with distances 2 and 16.
For the first query, between 0 and 1 there is no path where each distance is less than 2, thus we return false for this query.
For the second query, there is a path (0 -> 1 -> 2) of two edges with distances less than 5, thus we return true for this query.
Example 2:


Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
Output: [true,false]
Exaplanation: The above figure shows the given graph.


Constraints:

2 <= n <= 105
1 <= edgeList.length, queries.length <= 105
edgeList[i].length == 3
queries[j].length == 3
0 <= ui, vi, pj, qj <= n - 1
ui != vi
pj != qj
1 <= disi, limitj <= 109
There may be multiple edges between two nodes.
*/

//Функция distanceLimitedPathsExist принимает на вход число вершин n, список ребер edgeList и
//список запросов queries, и возвращает массив ans с булевыми значениями,
//где ans[i] равен true, если существует путь между вершинами queries[i][0] и queries[i][1] с ребрами длины меньше queries[i][2].
//
//Алгоритм состоит из следующих шагов:
//
//Для каждого запроса добавляем индекс запроса в конец массива.
//Сортируем массив запросов по возрастанию длины ребер limitj.
//Сортируем список ребер по возрастанию длины ребер disi.
//Инициализируем переменную i равной 0, которая будет использоваться для индексации в списке ребер.
//Для каждого запроса, начиная с самого короткого, соединяем все ребра

package main

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	uf := NewUnionFind(n)

	for i, query := range queries {
		queries[i] = append(query, i)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	i := 0 // i := edgeList's index
	for _, query := range queries {
		// Union edges whose distances < limit (query[2])
		for i < len(edgeList) && edgeList[i][2] < query[2] {
			uf.Union(edgeList[i][0], edgeList[i][1])
			i++
		}
		if uf.Find(query[0]) == uf.Find(query[1]) {
			ans[query[3]] = true
		}
	}

	return ans
}

type UnionFind struct {
	id   []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	id := make([]int, n)
	rank := make([]int, n)
	for i := range id {
		id[i] = i
	}
	return &UnionFind{id, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.id[x] == x {
		return x
	}
	uf.id[x] = uf.Find(uf.id[x])
	return uf.id[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.id[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.id[rootY] = rootX
	} else {
		uf.id[rootX] = rootY
		uf.rank[rootY]++
	}
}

func main() {
	n := 3
	edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries := [][]int{{0, 1, 2}, {0, 2, 5}}
	distanceLimitedPathsExist(n, edgeList, queries)
}

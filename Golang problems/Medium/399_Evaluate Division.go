/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.



Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
Example 2:

Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]
Example 3:

Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]
*/

package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1 / values[i]
	}
	results := make([]float64, len(queries))
	for i, query := range queries {
		a, b := query[0], query[1]
		if _, ok := graph[a]; !ok {
			results[i] = -1.0
			continue
		}
		if _, ok := graph[b]; !ok {
			results[i] = -1.0
			continue
		}
		if a == b {
			results[i] = 1.0
			continue
		}
		visited := make(map[string]bool)
		results[i] = dfs(graph, visited, a, b, 1.0)
	}
	return results
}

func dfs(graph map[string]map[string]float64, visited map[string]bool, curr string, target string, value float64) float64 {
	visited[curr] = true
	if _, ok := graph[curr]; !ok {
		return -1.0
	}
	if _, ok := graph[curr][target]; ok {
		return value * graph[curr][target]
	}
	for neighbor, weight := range graph[curr] {
		if !visited[neighbor] {
			if result := dfs(graph, visited, neighbor, target, value*weight); result != -1.0 {
				return result
			}
		}
	}
	return -1.0
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	calcEquation(equations, values, queries)
}

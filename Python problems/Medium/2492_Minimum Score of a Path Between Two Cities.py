# You are given a positive integer n representing n cities numbered from 1 to n. You are also given a 2D array roads where roads[i] = [ai, bi, distancei] indicates that there is a bidirectional road between cities ai and bi with a distance equal to distancei. The cities graph is not necessarily connected.
#
# The score of a path between two cities is defined as the minimum distance of a road in this path.
#
# Return the minimum possible score of a path between cities 1 and n.
#
# Note:
#
# A path is a sequence of roads between two cities.
# It is allowed for a path to contain the same road multiple times, and you can visit cities 1 and n multiple times along the path.
# The test cases are generated such that there is at least one path between 1 and n.
#
#
# Example 1:
#
#
# Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
# Output: 5
# Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5.
# It can be shown that no other path has less score.
# Example 2:
#
#
# Input: n = 4, roads = [[1,2,2],[1,3,4],[3,4,7]]
# Output: 2
# Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.
#
#
# Constraints:
#
# 2 <= n <= 105
# 1 <= roads.length <= 105
# roads[i].length == 3
# 1 <= ai, bi <= n
# ai != bi
# 1 <= distancei <= 104
# There are no repeated edges.
# There is at least one path between 1 and n.

# решение использует алгоритм обхода в ширину (BFS) для нахождения пути с минимальным значением.
#
# Кратко объясню ваш алгоритм:
#
# Инициализируйте переменную ans с бесконечностью. Создайте граф в виде списка смежности, где для каждого города u
# хранятся все города v, которые соединены с u, а также расстояние между ними. Начните BFS с города 1. Используйте
# очередь для хранения городов, которые нужно обработать, и множество seen для отслеживания уже посещенных городов.
# Переберите все города, которые связаны с текущим городом, и обновите переменную ans с минимальным расстоянием между
# текущим городом и городом v. Если город v не был посещен ранее, добавьте его в очередь и множество seen. Повторяйте
# шаги 4-5, пока очередь не пуста. Верните значение переменной ans. Время выполнения вашего алгоритма составляет O(n
# + m), где n - количество городов, а m - количество дорог в массиве roads. Это происходит из-за того, что каждая
# дорога обрабатывается только один раз.


import math
import collections
from typing import List


class Solution:
  def minScore(self, n: int, roads: List[List[int]]) -> int:
    ans = math.inf
    graph = [[] for _ in range(n + 1)]  # graph[u] := [(v, distance)]
    q = collections.deque([1])
    seen = {1}

    for u, v, distance in roads:
      graph[u].append((v, distance))
      graph[v].append((u, distance))

    while q:
      u = q.popleft()
      for v, d in graph[u]:
        ans = min(ans, d)
        if v in seen:
          continue
        q.append(v)
        seen.add(v)

    return ans


if __name__ == '__main__':
    n = 4
    roads = [[1, 2, 9], [2, 3, 6], [2, 4, 5], [1, 4, 7]]
    print(Solution().minScore(n, roads))


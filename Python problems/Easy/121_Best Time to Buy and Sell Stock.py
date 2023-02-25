# You are given an array prices where prices[i] is the price of a given stock on the ith day.
#
# You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
#
# Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
#
#
#
# Example 1:
#
# Input: prices = [7,1,5,3,6,4]
# Output: 5
# Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
# Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
# Example 2:
#
# Input: prices = [7,6,4,3,1]
# Output: 0
# Explanation: In this case, no transactions are done and the max profit = 0.
#
#
# Constraints:
#
# 1 <= prices.length <= 105
# 0 <= prices[i] <= 104
from typing import List


# Для решения данной задачи мы можем пройтись по массиву prices и отслеживать минимальное значение для покупки акций,
# а затем на каждой следующей итерации вычислять разницу между текущим значением цены и минимальным значением,
# чтобы узнать, является ли текущая цена оптимальной для продажи. Затем мы будем обновлять максимальную прибыль,
# если текущая прибыль больше максимальной.

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_price = float('inf')  # установим минимальную цену на покупку как бесконечность
        max_profit = 0  # установим максимальную прибыль как 0
        for price in prices:
            if price < min_price:
                min_price = price  # обновляем минимальную цену, если текущая цена меньше минимальной
            else:
                max_profit = max(max_profit,
                                 price - min_price)  # обновляем максимальную прибыль, если текущая прибыль больше
                # максимальной
        return max_profit


if __name__ == '__main__':
    print(Solution().maxProfit([7, 1, 5, 3, 6, 4]))
    print(Solution().maxProfit([7, 6, 4, 3, 1]))
    print(Solution().maxProfit([7, 6, 4, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))

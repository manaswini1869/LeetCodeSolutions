class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0
        left, right = 0, 1 # left buy right sell
        maxP = 0
        while right < len(prices):
            if prices[left] < prices[right]:
                pro = prices[right] - prices[left]
                maxP = max(maxP, pro)
            else:
                left = right
            right += 1
        return maxP


        
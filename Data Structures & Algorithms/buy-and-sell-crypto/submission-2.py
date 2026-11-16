class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minn = 100
        res = 0
        for price in prices:
            minn = min(minn, price)
            res = max(res, price-minn)
        return res
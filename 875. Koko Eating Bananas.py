class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        maxK = max(piles)
        minK = 1
        res = maxK
        while minK <= maxK:
            k = (minK + maxK) // 2
            hours = 0
            for i in piles:
                hours += math.ceil(i / k)
            if hours <= h:
                res = min(res, k)
                maxK = k-1
            else:
                minK = k + 1
        return res

            


        
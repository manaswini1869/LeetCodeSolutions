class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        numSet = set(nums)
        res = 0
        for i in nums:
            if (i-1) not in numSet:
                length = 0
                while (i+length) in numSet:
                    length += 1
                res = max(length, res)
        return res
            
        
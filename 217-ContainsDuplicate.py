class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        ele_count = {}
        for i in nums:
            if i in ele_count:
                ele_count[i] += 1
            else:
                ele_count[i] = 1
            if ele_count[i] >= 2:
                return True
        return False
        
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        n = len(nums)
        hash_res = {}
        for idx, val in enumerate(nums):
            req = target - val
            if req in hash_res:
                return [idx, hash_res[req]]
            else:
                hash_res[val] = idx
        
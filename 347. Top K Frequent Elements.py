class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        res_map = {}
        for num in nums:
            if num not in res_map:
                res_map[num] = 1
            res_map[num] += 1
        res_list = list(res_map.items())
        res_list.sort(key= lambda x:x[1], reverse= True)
        result = []
        n = 0
        print(res_list)
        while k > 0:
            result.append(res_list[n][0])
            n += 1
            k -= 1
        print(result)
        return result
        
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = {}
        for i in strs:
            sorted_s = "".join(sorted(i))
            if sorted_s in res:
                res[sorted_s].append(i)
            else:
                res[sorted_s] = []
                res[sorted_s].append(i)
        return res.values()
        
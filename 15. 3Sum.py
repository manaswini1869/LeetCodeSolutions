class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        n, p, z = [], [], []
        res = set()
        for num in nums:
            if num < 0:
                n.append(num)
            elif num > 0:
                p.append(num)
            else:
                z.append(num)
        N = set(n)
        P = set(p)
        if len(z)>=3:
            res.add((0, 0, 0))
        if z and N:
            for i in N:
                req = (-1)*(i)
                if req in P:
                    res.add(tuple(sorted((i,req, 0))))
        if z and P:
            for i in P:
                req = (-1)*(i)
                if req in N:
                    res.add(tuple(sorted((i,req, 0))))
        for i in range(len(p)):
            for j in range(i+1, len(p)):
                req = (-1)*(p[i]+p[j])
                if req in N:
                    res.add(tuple(sorted((p[i],p[j],req))))
        for i in range(len(n)):
            for j in range(i+1, len(n)):
                req = (-1)*(n[i]+n[j])
                if req in P:
                    res.add(tuple(sorted((n[i],n[j],req))))
        print(res)
        return res

        
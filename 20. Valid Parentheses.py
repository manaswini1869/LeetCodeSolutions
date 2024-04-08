class Solution:
    def isValid(self, s: str) -> bool:
        brac_map = {')':'(', '}':'{', ']':'['}
        res = []
        for i in s:
            if i in brac_map.values():
                res.append(i)
            else:
                if not res or res[-1] != brac_map[i]:
                    return False
                res.pop()
        if (len(res) > 0):
            return False
        else:
            return True
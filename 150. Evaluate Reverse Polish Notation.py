class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        res = []
        ops = ["+","-","*","/"]
        for token in tokens:
            if token in ops and res:
                a = res.pop()
                b = res.pop()
                if token == ops[0]:
                    res.append(a+b)
                if token == ops[1]:
                    res.append(b-a)
                if token == ops[2]:
                    res.append(a*b)
                if token == ops[3]:
                    res.append(int(b/a))
            else:
                res.append(int(token))
        return int(res[0])
        
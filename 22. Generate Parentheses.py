class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        # add open only if open count > close count
        # add close only if close count < open count
        # till n == open == close

        stack = [] # hold parans
        res = []

        def backtrack(openn, closen):
            if openn == closen == n:
                res.append("".join(stack))
                return
            if openn < n:
                stack.append('(')
                backtrack(openn+1, closen)
                stack.pop()
            if closen < openn:
                stack.append(')')
                backtrack(openn, closen+1)
                stack.pop()
        backtrack(0,0)
        return res

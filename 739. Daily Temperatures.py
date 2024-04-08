class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        answer = [0]*len(temperatures)
        stack = []

        for idx, temp in enumerate(temperatures):
            while stack and temp > stack[-1][0]:
                prevt, prevI = stack.pop()
                answer[prevI] = (idx - prevI)
            stack.append([temp,idx])
        print(answer)
        return answer



        
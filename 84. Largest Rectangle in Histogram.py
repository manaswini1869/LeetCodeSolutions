class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = [] #pair index, height
        area = 0
        for idx, hght in enumerate(heights):
            start = idx
            while stack and stack[-1][1] > hght:
                index, height = stack.pop()
                area = max(area, height*(idx-index))
                start = index
            stack.append([start, hght])
        for i, h in stack:
            area = max(area, (h*(len(heights)-i)))
        print(stack)
        return area

        
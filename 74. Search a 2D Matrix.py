class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        n = len(matrix)
        m = len(matrix[0])
        row = 0
        for i in range(n):
            if target >= matrix[i][0] and target <= matrix[i][m-1]:
                row = i
        print(row)
        arr = matrix[row]
        print(arr)
        left, right = 0, len(arr)-1
        while left <= right:
            mid = (left + right) // 2
            if target < arr[mid]:
                right -= 1
            elif target > arr[mid]:
                left += 1
            else:
                return True
        return False
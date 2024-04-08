class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        right, left = 0, len(numbers)-1
        def calSum(right, left):
            addition = numbers[right] + numbers[left]
            if addition > target:
                left -= 1
            elif addition < target:
                right += 1
            else:
                return [right+1, left+1]
            return calSum(right, left)
        return calSum(right, left)
        
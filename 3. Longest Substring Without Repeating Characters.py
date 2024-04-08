class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        if len(s) == 1:
            return 1
        if len(s) > 1:
            right, left = 0, 0
            maxL = 0
            ans = set()
            while right < len(s) and left < len(s):
                if s[left] in ans:
                    ans.remove(s[right])
                    right += 1
                else:
                    ans.add(s[left])
                    left += 1
                    maxL = max(maxL, len(ans))
        return maxL
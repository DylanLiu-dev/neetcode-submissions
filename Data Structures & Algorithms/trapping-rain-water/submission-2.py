class Solution:
    def trap(self, height: List[int]) -> int:
       # 1. pre-calculate the height of wall of left side and right side
       # 2. sum trapping water in it
        left, right = [], []
        wall = 0
        for h in height:
            wall = max(wall, h)
            left.append(wall)
        wall = 0
        for i in range(len(height)-1, -1, -1):
            h = height[i]
            wall = max(wall, h)
            right.append(wall)
        right.reverse()
        res = 0
        for i, h in enumerate(height):
            res += min(left[i], right[i])-h
        return res
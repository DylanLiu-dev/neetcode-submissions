class Solution:
    def trap(self, height: List[int]) -> int:
       # 1. pre-calculate the height of wall of left side and right side
       # 2. sum trapping water in it
        left, right = [0 for _ in range(len(height))], [0 for _ in range(len(height))]
        wall = 0
        for i, h in enumerate(height):
            wall = max(wall, h)
            left[i] = wall
        wall = 0
        for i in range(len(height)-1, -1, -1):
            h = height[i]
            wall = max(wall, h)
            right[i] = wall
        res = 0
        for i, h in enumerate(height):
            res += min(left[i], right[i])-h
        return res
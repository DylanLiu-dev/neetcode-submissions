class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        n = len(nums)
        mp = {}
        for i in range(n):
            complement = target - nums[i]
            if complement in mp:
                return [mp[complement], i]
            mp[nums[i]] = i
        return []
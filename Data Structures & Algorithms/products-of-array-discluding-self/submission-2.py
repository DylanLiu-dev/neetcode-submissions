class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        # pre calculate
        # 1   2   4   6
        # 1   2   8  24   leftp
        #48  48  24   6   rightp
        #48  24  12   8
        leftp, rightp = [1 for i in range(len(nums))], [1 for i in range(len(nums))]
        for i in range(len(nums)):
            leftp[i] *= nums[i]
            if i-1 >= 0:
                leftp[i] *= leftp[i-1]
        for i in range(len(nums)-1, -1, -1):
            rightp[i] *= nums[i]
            if i+1 < len(nums):
                rightp[i] *= rightp[i+1]
        res = [1 for i in range(len(nums))]
        for i in range(len(nums)):
            if i-1 < 0:
                res[i] = rightp[i+1]
                continue
            if i+1 >= len(nums):
                res[i] = leftp[i-1]
                continue
            res[i] *= leftp[i-1] * rightp[i+1]
        return res
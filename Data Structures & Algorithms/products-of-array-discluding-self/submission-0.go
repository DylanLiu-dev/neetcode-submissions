func productExceptSelf(nums []int) []int {
	presumLtoR := make([]int, len(nums))
	presumRtoL := make([]int, len(nums))
	for i := range nums{
		presumLtoR[i] = nums[i]
		presumRtoL[i] = nums[i]
	}
	for i:=1;i<len(presumLtoR); i++{
		presumLtoR[i] *= presumLtoR[i-1]
	}
	for i:=len(presumRtoL)-2;i>-1; i--{
		presumRtoL[i] *= presumRtoL[i+1]
	}
	res := make([]int, 0, len(nums))
	for i := range nums{
		if i == 0 {
			res = append(res, presumRtoL[i+1])
		}else if i == len(nums)-1 {
			res = append(res, presumLtoR[i-1])
		}else{
			res = append(res, presumRtoL[i+1]*presumLtoR[i-1])
		}
	}
	return res
}

// [ 1, 2, 4, 6]
// [ 1, 2, 8,48]  presumLtoR
// [48,48,24, 6]  presumRtoL
//  48,24,12, 8

// [-1,0,1,2,3]
// [-1,0,0,0,0]
// [ 0,0,6,6,3]
// [ 0,-6,0,0,0]
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i:=0; i<len(nums)-2; i++ {
		val := nums[i]
		if val > 0 {
			break
		}
		if i > 0 && val == nums[i-1]{
			// skil duplicate val
			continue
		}
		l, r := i+1, len(nums)-1
		for l<r {
			sum := val + nums[l] + nums[r]
			if sum > 0 {
				r--
			}else if sum < 0 {
				l++
			}else {
				res = append(res, []int{val, nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l-1] {
					// skip duplicate nums[l]
					l++
				}
			}
		}
	}
	return res
}

// 1. sort it
// 2. apply Two Sums II

func longestConsecutive(nums []int) int {
	track := make(map[int]int)
	res := 0
	for _, num := range nums {
		if track[num] == 0 {
			left := track[num-1]
			right := track[num+1]
			sum := left+right+1
			track[num] = sum
			track[num-left] =sum
			track[num+right] = sum
			if sum > res {
				res = sum
			}
		}
	}
	return res
}

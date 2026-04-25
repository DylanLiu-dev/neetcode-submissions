func twoSum(nums []int, target int) []int {
    idxMap := map[int]int{}
	for idx, num := range nums {
		if j, exist := idxMap[target-num]; exist {
			return []int{j, idx}
		}
		idxMap[num] = idx
	}
	return nil
}

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	pairs := make([][]int, 0, len(cnt))
	for n, f := range cnt {
		pairs = append(pairs, []int{f, n})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})
	var res []int
	for i := 0; i<k; i++ {
		res = append(res, pairs[i][1])
	}
	return res
}

func maxArea(heights []int) int {
	var res int
	i, j := 0, len(heights)-1
	for i<j {
		if heights[i] < heights[j] {
			res = max(res, (j-i)*heights[i])
			i++
			continue
		}
		res = max(res, (j-i)*heights[j])
		j--
	}
	return res
}

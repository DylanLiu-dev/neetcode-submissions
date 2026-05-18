type s struct{
	temp int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []s{}
	for i, t := range temperatures {
		for len(stack) > 0 && t > stack[len(stack)-1].temp {
			stackIdx := stack[len(stack)-1].idx
			res[stackIdx] = i - stackIdx
			stack = stack[0:len(stack)-1]
		}
		stack = append(stack, s{temp: t, idx: i})
	}
	return res
}
// monotonic stack
// 40 36
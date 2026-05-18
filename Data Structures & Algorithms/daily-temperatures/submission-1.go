func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stk := []int{} // the stack of days that need to know when will get warmer
	for i, t := range temperatures {
		for len(stk) > 0 && t > temperatures[stk[len(stk)-1]] {
			var top int
			top, stk = stk[len(stk)-1], stk[:len(stk)-1]
			res[top] = i - top
		}
		stk = append(stk, i)
	}
	return res
}
// monotonic stack
// 40 36
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	pairs := make([][2]int, n)
	for i:=0; i<n; i++ {
		pairs[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})
	stk := []float64{}
	for _, p := range pairs {
		time := float64(target-p[0])/float64(p[1])
		stk = append(stk, time)
		if len(stk) >= 2 && stk[len(stk)-1] <= stk[len(stk)-2] {
			stk = stk[:len(stk)-1]
		} 
	}
	return len(stk)
}

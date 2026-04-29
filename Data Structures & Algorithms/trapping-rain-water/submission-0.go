func trap(height []int) int {
	prehighest := make([]int, len(height))
	posthighest := make([]int, len(height))
	prehighest[0] = height[0]
	for i:=1; i<len(height); i++ {
		prehighest[i] = max(prehighest[i-1], height[i])
	}
	posthighest[len(height)-1] = height[len(height)-1]
	for i:=len(height)-2; i>=0; i-- {
		posthighest[i] = max(height[i], posthighest[i+1])
	}
	var res int
	for i:=0; i<len(height); i++ {
		res += min(prehighest[i], posthighest[i]) - height[i]
	}
	return res
}

// 1. calculate the prehighest and posthighest arrays
// 2. sum of min(prehighest[i]+posthighest[i])-height[i]
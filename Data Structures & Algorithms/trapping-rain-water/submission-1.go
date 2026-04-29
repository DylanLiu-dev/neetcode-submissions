func trap(height []int) int {
	n := len(height)
	prehighest := make([]int, n)
	posthighest := make([]int, n)
	prehighest[0] = height[0]
	for i:=1; i<n; i++ {
		prehighest[i] = max(prehighest[i-1], height[i])
	}
	posthighest[n-1] = height[n-1]
	for i:=n-2; i>=0; i-- {
		posthighest[i] = max(height[i], posthighest[i+1])
	}
	var res int
	for i:=0; i<n; i++ {
		res += min(prehighest[i], posthighest[i]) - height[i]
	}
	return res
}

// 1. calculate the prehighest and posthighest arrays
// 2. sum of min(prehighest[i]+posthighest[i])-height[i]
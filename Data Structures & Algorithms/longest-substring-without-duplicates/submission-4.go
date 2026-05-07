func lengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	l, res := 0, 0
	for r:=0; r<len(s); r++ {
		if idx, ok := seen[s[r]]; ok {
			l = max(idx+1, l)
		}
		seen[s[r]] = r
		if r-l+1 > res {
			res = r-l+1
		}
	}
	return res
}
// 0. consider a sliding window which left index is i, right is j
// 1. extend the window until meet repetitive
// 2. shrink the window and eliminate the element pointed by i
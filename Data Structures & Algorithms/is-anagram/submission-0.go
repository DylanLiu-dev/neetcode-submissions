func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, c := range s {
		cnt[c] += 1
	}
	for _, c := range t {
		cnt[c] -= 1
	}
	for _, n := range cnt {
		if n != 0 {
			return false
		}
	}
	return true
}

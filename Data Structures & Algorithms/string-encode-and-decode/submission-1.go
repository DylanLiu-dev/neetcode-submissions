type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, str := range strs {
		res += strconv.Itoa(len(str))
		res += "#"
		res += str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	i := 0
	for i<len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		n, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+n])
		i += n
	}
	return res
}

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
		var tmp string
		for string(encoded[i]) != "#" {
			tmp += string(encoded[i])
			i++
		}
		n, _ := strconv.Atoi(tmp)
		i++
		res = append(res, encoded[i:i+n])
		i += n
	}
	return res
}

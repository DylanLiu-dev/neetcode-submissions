func isValid(s string) bool {
    stack := make([]rune, 0)
	for _, c := range s {
		switch c {
			case '(', '{', '[':
			stack = append(stack, c)
			case ')':
			n := len(stack) - 1
			if n == -1 || stack[n] != '(' {
				return false
			}
			stack = stack[:n]
			case '}':
			n := len(stack) - 1
			if n==-1||stack[n] != '{' {
				return false
			}
			stack = stack[:n]
			case ']':
			n := len(stack) - 1
			if n == -1 ||stack[n] != '[' {
				return false
			}
			stack = stack[:n]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		n := len(stack)
		switch token {
			case "+":
			a, b := stack[n-2], stack[n-1]
			stack = stack[0:n-2]
			stack = append(stack, a+b)
			case "-":
			a, b := stack[n-2], stack[n-1]
			stack = stack[0:n-2]
			stack = append(stack, a-b)
			case "*":
			a, b := stack[n-2], stack[n-1]
			stack = stack[0:n-2]
			stack = append(stack, a*b)
			case "/":
			a, b := stack[n-2], stack[n-1]
			stack = stack[0:n-2]
			stack = append(stack, a/b)
			default:
			v, _ := strconv.Atoi(token)
			stack = append(stack, v)
		}
	}
	return stack[0]
}

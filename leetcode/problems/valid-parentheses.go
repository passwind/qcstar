package problems

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	str := []byte(s)
	stack := make([]byte, len(s))
	i := 0
	for _, c := range str {
		if c == '(' || c == '[' || c == '{' {
			stack[i] = c
			i++
		} else if i > 0 && ((stack[i-1] == '(' && c == ')') || (stack[i-1] == '[' && c == ']') || (stack[i-1] == '{' && c == '}')) {
			i--
		} else {
			return false
		}
	}
	return i == 0
}

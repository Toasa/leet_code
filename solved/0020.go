func openParen(c byte) bool {
	return (c == '(') || (c == '{') || (c == '[')
}

func match(open, close byte) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	}
	return false
}

func isValid(s string) bool {
	stack := []byte{}

	if len(s) == 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if openParen(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack)-1 < 0 {
				return false
			}
			if match(stack[len(stack)-1], s[i]) {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
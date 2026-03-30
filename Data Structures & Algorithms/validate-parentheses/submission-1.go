func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		// add to stack
		case '{', '[', '(':
			stack = append(stack, c)

		// pop from stack and eval
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}

			// comparator 
			var comparator rune
			switch c {
				case '}':
				comparator = '{'
				case ']':
				comparator = '['
				case ')':
				comparator = '('
			}

			var val rune
			val, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if val != comparator {
				return false
			}
		}
	}

	return len(stack) == 0
}

package stackqueue

// IsBalanced checks if the parentheses in a string are balanced
func IsBalanced(input string) bool {
	stack := Stack{}

	for _, char := range input {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		} else if char == ')' || char == '}' || char == ']' {
			if stack.IsEmpty() {
				return false
			}
			top, _ := stack.Pop()
			if (char == ')' && top != '(') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

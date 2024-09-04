package stackqueue

import (
	"errors"
	"strconv"
)

// EvaluatePostfix evaluates a postfix expression
func EvaluatePostfix(expression string) (int, error) {
	stack := Stack{}

	for _, char := range expression {
		if char >= '0' && char <= '9' {
			value, _ := strconv.Atoi(string(char))
			stack.Push(rune(value))
		} else {
			operand2, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			operand1, err := stack.Pop()
			if err != nil {
				return 0, err
			}

			var result int
			switch char {
			case '+':
				result = int(operand1) + int(operand2)
			case '-':
				result = int(operand1) - int(operand2)
			case '*':
				result = int(operand1) * int(operand2)
			case '/':
				result = int(operand1) / int(operand2)
			default:
				return 0, errors.New("invalid operator")
			}

			stack.Push(rune(result))
		}
	}

	result, _ := stack.Pop()
	return int(result), nil
}

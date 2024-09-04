package stackqueue

import (
	"fmt"
	"testing"
)

func TestEvaluatePostfix(t *testing.T) {
	expression := "231*+9-"
	result, err := EvaluatePostfix(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The result of the postfix expression %s is: %d\n", expression, result)
	}
}

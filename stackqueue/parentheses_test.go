package stackqueue

import (
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	expression := "{[()()]}"
	fmt.Printf("Is the expression %s balanced? %v\n", expression, IsBalanced(expression))
}

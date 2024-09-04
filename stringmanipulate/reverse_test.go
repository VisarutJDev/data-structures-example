package stringmanipulate

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	str := "hello"
	reversed := ReverseString(str)
	fmt.Println("Reversed string:", reversed)
}

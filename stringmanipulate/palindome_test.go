package stringmanipulate

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	str := "madam"
	if IsPalindrome(str) {
		fmt.Printf("%s is a palindrome\n", str)
	} else {
		fmt.Printf("%s is not a palindrome\n", str)
	}
}

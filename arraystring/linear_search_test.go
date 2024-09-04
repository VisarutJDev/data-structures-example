package arraystring

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{4, 2, 7, 1, 9}
	target := 7
	index := LinearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in array\n", target)
	}
}

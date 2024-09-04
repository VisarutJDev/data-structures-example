package mysort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	Bubble(arr)
	fmt.Println("Sorted array:", arr)
}

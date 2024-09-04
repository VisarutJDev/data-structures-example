package mysort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", arr)

	InsertionSort(arr)
	fmt.Println("Sorted array:", arr)
}

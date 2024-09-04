package mysort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	sortedArr := MergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

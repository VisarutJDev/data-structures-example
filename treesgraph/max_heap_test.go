package treesgraph

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := &MaxHeap{}
	fmt.Println("Inserting elements into the heap:")
	h.Insert(3)
	h.Insert(1)
	h.Insert(6)
	h.Insert(5)
	h.Insert(2)
	h.Insert(4)

	fmt.Println("Max-Heap array:", h.array)

	fmt.Println("Extracting elements from the heap:")
	fmt.Println(h.ExtractMax()) // Output: 6
	fmt.Println(h.ExtractMax()) // Output: 5
	fmt.Println(h.ExtractMax()) // Output: 4
	fmt.Println(h.ExtractMax()) // Output: 3
	fmt.Println(h.ExtractMax()) // Output: 2
	fmt.Println(h.ExtractMax()) // Output: 1
}

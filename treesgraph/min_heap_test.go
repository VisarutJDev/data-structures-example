package treesgraph

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &MinHeap{}
	fmt.Println("Inserting elements into the heap:")
	h.Insert(3)
	fmt.Println("Min-Heap array:", h.array)
	h.Insert(1)
	fmt.Println("Min-Heap array:", h.array)
	h.Insert(6)
	fmt.Println("Min-Heap array:", h.array)
	h.Insert(5)
	fmt.Println("Min-Heap array:", h.array)
	h.Insert(2)
	fmt.Println("Min-Heap array:", h.array)
	h.Insert(4)

	fmt.Println("Min-Heap array:", h.array)

	fmt.Println("Extracting elements from the heap:")
	fmt.Println(h.ExtractMin()) // Output: 1
	fmt.Println("Min-Heap array:", h.array)
	fmt.Println(h.ExtractMin()) // Output: 2
	fmt.Println("Min-Heap array:", h.array)
	fmt.Println(h.ExtractMin()) // Output: 3
	fmt.Println("Min-Heap array:", h.array)
	fmt.Println(h.ExtractMin()) // Output: 4
	fmt.Println("Min-Heap array:", h.array)
	fmt.Println(h.ExtractMin()) // Output: 5
	fmt.Println("Min-Heap array:", h.array)
	fmt.Println(h.ExtractMin()) // Output: 6
	fmt.Println("Min-Heap array:", h.array)
}

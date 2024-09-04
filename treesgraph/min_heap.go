package treesgraph

import (
	"fmt"
)

// MinHeap represents a Min-Heap data structure
type MinHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// ExtractMin removes and returns the smallest element from the heap
func (h *MinHeap) ExtractMin() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	min := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	fmt.Println("h.array[:lastIndex]", h.array)
	h.heapifyDown(0)
	return min
}

// heapifyUp restores the heap property by moving the element at index i up
func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown restores the heap property by moving the element at index i down
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// Loop until index is not the last index and any child is smaller
	for l <= lastIndex {
		if l == lastIndex { // When left child is the only child
			childToCompare = l
		} else if h.array[l] < h.array[r] { // When left child is smaller
			childToCompare = l
		} else { // When right child is smaller
			childToCompare = r
		}

		// If the parent is smaller than the smallest child, break
		if h.array[index] < h.array[childToCompare] {
			break
		}

		// Swap and continue
		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
		fmt.Println("Min-Heap heapifyDown:", h.array)
	}
}

// parent returns the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the left child index
func left(i int) int {
	return 2*i + 1
}

// right returns the right child index
func right(i int) int {
	return 2*i + 2
}

// swap swaps the elements at indexes i and j
func (h *MinHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

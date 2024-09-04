package linkedlists

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create a new linked list
	ll := &LinkedList{}

	// Insert elements into the linked list
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.Print() // Output: 1 -> 2 -> 3 -> 4 -> nil

	// Delete an element
	ll.Delete(3)
	ll.Print() // Output: 1 -> 2 -> 4 -> nil

	// Reverse the linked list
	ll.Reverse()
	ll.Print() // Output: 4 -> 2 -> 1 -> nil

	// Create a cycle for testing
	ll.Head.Next.Next.Next = ll.Head.Next // Create a cycle: 4 -> 2 -> 1 -> 2

	// Detect cycle
	hasCycle := ll.HasCycle()
	fmt.Printf("Cycle detected: %v\n", hasCycle) // Output: Cycle detected: true
}

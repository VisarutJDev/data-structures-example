package linkedlists

import (
	"fmt"
)

// Node represents a node in a singly linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
}

// InsertAtEnd inserts a new node at the end of the linked list
func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Delete deletes a node with the given value from the linked list
func (ll *LinkedList) Delete(value int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Reverse reverses the linked list
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}

// HasCycle detects if there is a cycle in the linked list
// Uses Floyd's Cycle-Finding Algorithm (Tortoise and Hare) to detect if there is a cycle in the linked list.
func (ll *LinkedList) HasCycle() bool {
	if ll.Head == nil {
		return false
	}
	slow := ll.Head
	fast := ll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Print displays the linked list
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

package stackqueue

import (
	"errors"
)

// Stack represents a stack data structure
type Stack struct {
	items []rune
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (rune, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

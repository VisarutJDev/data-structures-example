package stackqueue

import (
	"errors"
)

// Queue represents a queue data structure
type Queue struct {
	items []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

package greedy

import (
	"fmt"
	"testing"
)

func TestActivitySelection(t *testing.T) {
	activities := []Activity{
		{start: 1, end: 4},
		{start: 3, end: 5},
		{start: 0, end: 6},
		{start: 5, end: 7},
		{start: 8, end: 9},
		{start: 5, end: 9},
	}

	selected := ActivitySelection(activities)
	fmt.Println("Selected activities:")
	for _, activity := range selected {
		fmt.Printf("Start: %d, End: %d\n", activity.start, activity.end)
	}
}

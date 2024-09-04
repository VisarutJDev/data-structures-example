package treesgraph

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	// Create a sample binary tree:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Right = NewNode(6)

	fmt.Println("Breadth First Search (BFS):")
	BFS(root) // Output: 1 2 3 4 5 6
}

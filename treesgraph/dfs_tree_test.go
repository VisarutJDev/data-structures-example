package treesgraph

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
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

	fmt.Println("Depth First Search (DFS):")
	DFS(root) // Output: 1 2 4 5 3 6
}

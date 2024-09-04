package treesgraph

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	fmt.Print("Pre-order: ")
	PreOrderTraversal(root) // Output: 1 2 4 5 3
	fmt.Println()

	fmt.Print("In-order: ")
	InOrderTraversal(root) // Output: 4 2 5 1 3
	fmt.Println()

	fmt.Print("Post-order: ")
	PostOrderTraversal(root) // Output: 4 5 2 3 1
	fmt.Println()
}

package treesgraph

import "fmt"

// NewNode creates a new TreeNode
func NewNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// DFS performs a depth-first search on the binary tree
func DFS(node *TreeNode) {
	if node == nil {
		return
	}

	// Process the current node (Pre-order)
	fmt.Printf("%d ", node.Value)

	// Recursively traverse the left subtree
	DFS(node.Left)

	// Recursively traverse the right subtree
	DFS(node.Right)
}

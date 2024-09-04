package treesgraph

import (
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraversal traverses the tree in pre-order (Root -> Left -> Right)
func PreOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	PreOrderTraversal(node.Left)
	PreOrderTraversal(node.Right)
}

// InOrderTraversal traverses the tree in in-order (Left -> Root -> Right)
func InOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	InOrderTraversal(node.Left)
	fmt.Printf("%d ", node.Value)
	InOrderTraversal(node.Right)
}

// PostOrderTraversal traverses the tree in post-order (Left -> Right -> Root)
func PostOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrderTraversal(node.Left)
	PostOrderTraversal(node.Right)
	fmt.Printf("%d ", node.Value)
}

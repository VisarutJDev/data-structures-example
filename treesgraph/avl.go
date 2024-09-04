package treesgraph

import (
	"fmt"
)

// AVLNode represents a node in an AVL tree
type AVLNode struct {
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// GetHeight returns the height of the node
func (node *AVLNode) GetHeight() int {
	if node == nil {
		return 0
	}
	return node.Height
}

// GetBalanceFactor returns the balance factor of the node
func (node *AVLNode) GetBalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.Left.GetHeight() - node.Right.GetHeight()
}

// RotateRight performs a right rotation on the subtree rooted at node
func RotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(y.Left.GetHeight(), y.Right.GetHeight()) + 1
	x.Height = max(x.Left.GetHeight(), x.Right.GetHeight()) + 1

	// Return new root
	return x
}

// RotateLeft performs a left rotation on the subtree rooted at node
func RotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(x.Left.GetHeight(), x.Right.GetHeight()) + 1
	y.Height = max(y.Left.GetHeight(), y.Right.GetHeight()) + 1

	// Return new root
	return y
}

// Insert inserts a value into the AVL tree and balances it
func (node *AVLNode) Insert(value int) *AVLNode {
	if node == nil {
		return &AVLNode{Value: value, Height: 1}
	}
	if value < node.Value {
		node.Left = node.Left.Insert(value)
	} else if value > node.Value {
		node.Right = node.Right.Insert(value)
	} else {
		return node // Duplicate values are not allowed
	}

	// Update the height of the current node
	node.Height = 1 + max(node.Left.GetHeight(), node.Right.GetHeight())

	// Get the balance factor
	balanceFactor := node.GetBalanceFactor()

	// Left Left Case
	if balanceFactor > 1 && value < node.Left.Value {
		return RotateRight(node)
	}

	// Right Right Case
	if balanceFactor < -1 && value > node.Right.Value {
		return RotateLeft(node)
	}

	// Left Right Case
	if balanceFactor > 1 && value > node.Left.Value {
		node.Left = RotateLeft(node.Left)
		return RotateRight(node)
	}

	// Right Left Case
	if balanceFactor < -1 && value < node.Right.Value {
		node.Right = RotateRight(node.Right)
		return RotateLeft(node)
	}

	// Return the (unchanged) node pointer
	return node
}

// AVLInOrderTraversal traverses the tree in in-order (Left -> Root -> Right)
func AVLInOrderTraversal(node *AVLNode) {
	if node == nil {
		return
	}
	AVLInOrderTraversal(node.Left)
	fmt.Printf("%d ", node.Value)
	AVLInOrderTraversal(node.Right)
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

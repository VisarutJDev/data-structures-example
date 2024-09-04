package treesgraph

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	root := &AVLNode{Value: 10, Height: 1}
	root = root.Insert(20)
	AVLInOrderTraversal(root)
	fmt.Println()
	root = root.Insert(30)
	AVLInOrderTraversal(root)
	fmt.Println()
	root = root.Insert(40)
	AVLInOrderTraversal(root)
	fmt.Println()
	root = root.Insert(50)
	AVLInOrderTraversal(root)
	fmt.Println()
	root = root.Insert(25)
	AVLInOrderTraversal(root)
	fmt.Println()

	fmt.Print("In-order traversal of the AVL tree: ")
	AVLInOrderTraversal(root) // Output: 10 20 25 30 40 50
	fmt.Println()
}

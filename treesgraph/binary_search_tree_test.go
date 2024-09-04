package treesgraph

import (
	"fmt"
	"testing"
)

func TestSearchTree(t *testing.T) {
	root := &BSTNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)

	fmt.Println("Search 7:", root.Search(7))   // Output: true
	fmt.Println("Search 12:", root.Search(12)) // Output: false
}

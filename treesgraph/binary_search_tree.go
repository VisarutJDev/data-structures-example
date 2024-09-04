package treesgraph

// BSTNode represents a node in a binary search tree
type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

// Insert inserts a value into the binary search tree
func (node *BSTNode) Insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BSTNode{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BSTNode{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

// Search searches for a value in the binary search tree
func (node *BSTNode) Search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.Search(value)
	} else if value > node.Value {
		return node.Right.Search(value)
	}
	return true
}

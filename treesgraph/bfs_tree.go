package treesgraph

import "fmt"

// BFS performs a breadth-first search on the binary tree
func BFS(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Process the current node
		fmt.Printf("%d ", current.Value)

		// Enqueue left child
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		// Enqueue right child
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

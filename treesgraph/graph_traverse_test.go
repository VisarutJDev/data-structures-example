package treesgraph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	fmt.Print("DFS starting from vertex 0: ")
	visited := make(map[int]bool)
	g.DFS(0, visited) // Output: 0 1 2 3 4
	fmt.Println()

	fmt.Print("BFS starting from vertex 0: ")
	g.BFS(0) // Output: 0 1 2 3 4
	fmt.Println()
}

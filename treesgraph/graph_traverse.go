package treesgraph

import (
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v) // For an undirected graph
}

// DFS performs Depth-First Search on the graph
func (g *Graph) DFS(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Printf("%d ", v)

	for _, neighbor := range g.vertices[v] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

// BFS performs Breadth-First Search on the graph
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", v)

		for _, neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	nodes map[string][]string
}

func (g *Graph) addEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
	g.nodes[to] = append(g.nodes[to], from) // For undirected graph
}

func (g *Graph) BFS(start, end string) []string {
	queue := list.New()
	queue.PushBack(start)
	visited := make(map[string]bool)
	parent := make(map[string]string)

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(string)
		visited[current] = true

		for _, neighbor := range g.nodes[current] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				parent[neighbor] = current
				if neighbor == end {
					return buildPath(parent, start, end)
				}
			}
		}
	}
	return nil // No path found
}

func buildPath(parent map[string]string, start, end string) []string {
	path := []string{end}
	for current := end; current != start; current = parent[current] {
		path = append([]string{parent[current]}, path...)
	}
	return path
}

func main() {
	g := &Graph{nodes: make(map[string][]string)}
	g.addEdge("A", "B")
	g.addEdge("A", "C")
	g.addEdge("B", "D")
	g.addEdge("C", "D")
	g.addEdge("C", "E")
	g.addEdge("D", "E")

	path := g.BFS("A", "E")
	fmt.Println(path) // Output: [A C E]
}

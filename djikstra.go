package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Define a struct to represent an edge in the graph
type Edge struct {
	node   string
	weight int
}

// Define a struct to represent the graph
type Graph map[string][]Edge

// Define a struct for the priority queue
type Item struct {
	node     string
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)

	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0
	heap.Push(pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.node

		if visited[node] {
			continue
		}
		visited[node] = true

		for _, edge := range graph[node] {
			neighbor := edge.node
			newDistance := distances[node] + edge.weight
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				heap.Push(pq, &Item{node: neighbor, distance: newDistance})
			}
		}
	}

	return distances
}

func main() {
	graph := Graph{
		"A": []Edge{{"B", 1}, {"C", 2}},
		"B": []Edge{{"D", 2}, {"E", 4}},
		"C": []Edge{{"B", 1}, {"D", 4}, {"E", 5}},
		"D": []Edge{{"E", 1}},
		"E": []Edge{},
	}

	distances := Dijkstra(graph, "A")
	fmt.Println(distances)
}

package dijkstra

import (
	"fmt"
	"golearning/algorithms/minHeapTuple"
	"log"
)

type Node struct {
	neighbour string
	wieght    int
}

var graph = make(map[string][]Node)

var distance = make(map[string]int)

func Run() {
	log.Println("DFS Struct")
	A := []Node{}
	B := []Node{}
	C := []Node{}
	D := []Node{}
	E := []Node{}

	// // Define the edges
	D = append(D, Node{neighbour: "E", wieght: 5})
	C = append(C, Node{neighbour: "E", wieght: 1})
	B = append(B, Node{neighbour: "D", wieght: 2})
	A = append(A, Node{neighbour: "B", wieght: 1})
	A = append(A, Node{neighbour: "C", wieght: 4})

	graph["A"] = A
	graph["B"] = B
	graph["C"] = C
	graph["D"] = D
	graph["E"] = E

	fmt.Println("graph", graph)
	for index := range graph {
		distance[index] = 9999999
	}
	distance["A"] = 0
	minHeap := minHeapTuple.Heap
	minHeap = minHeapTuple.Insert(minHeapTuple.HeapNode{Priority: 0, Value: "A"}, minHeap)
	fmt.Println("Heap after Insertion:", minHeap)

	var visited = make(map[string]bool)

	for len(minHeap) > 0 {
		current_distance := minHeap[0].Priority
		current_node := minHeap[0].Value
		minHeap = minHeapTuple.Remove(minHeap)
		fmt.Println("current_distance: ", current_distance)
		fmt.Println("current_node: ", current_node)

		if visited[current_node] {
			continue
		}
		visited[current_node] = true

		for _, node := range graph[current_node] {
			neighbour := node.neighbour
			weight := node.wieght
			distance_through_current := current_distance + weight

			if distance_through_current < distance[neighbour] {
				distance[neighbour] = distance_through_current
				minHeap = minHeapTuple.Insert(minHeapTuple.HeapNode{Priority: distance_through_current, Value: neighbour}, minHeap)
			}

		}
	}
	fmt.Println("distance", distance)
}

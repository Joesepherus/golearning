package bfsStruct

import (
	"golearning/queue"
	"log"
)

type Node struct {
	node  string
	edges []Node
}

var visited = make(map[string]bool)

var graph Graph

type Graph struct {
	nodes []Node
}

var q = &queue.Queue{}

func Run() {
	log.Println("BFS Struct")
	A := Node{node: "A", edges: []Node{}}
	B := Node{node: "B", edges: []Node{}}
	C := Node{node: "C", edges: []Node{}}
	D := Node{node: "D", edges: []Node{}}
	E := Node{node: "E", edges: []Node{}}
	F := Node{node: "F", edges: []Node{}}

	E.edges = []Node{F}
	C.edges = []Node{}
	B.edges = []Node{D, E}
	A.edges = []Node{B, C}

	graph = Graph{
		nodes: []Node{A, B, C, D, E, F},
	}

	log.Println("q", q)
	BFS(A)
}

func BFS(start Node) {
	node := start
	q.Enqueue(node)
	log.Println("START NODE adding to queue: ", node)

	for !q.IsEmpty() {
		nodeAny, _ := q.Dequeue() // Dequeue returns an `any`

		node = nodeAny.(Node)
		println("traversing node ", node.node)

		for _, edge := range node.edges {
			log.Println("adding to queue: ", edge)
			q.Enqueue(edge)
		}

	}
}

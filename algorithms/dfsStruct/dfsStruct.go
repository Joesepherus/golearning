package dfsStruct

import "log"

type Node struct {
	node  string
	edges []Node
}

var visited = make(map[string]bool)

// var graph = map[string][]string{
// 	"A": {"B", "C"},
// 	"B": {"D", "E"},
// 	"C": {"F"},
// 	"D": {},
// 	"E": {"F"},
// 	"F": {},
// }

var graph Graph

type Graph struct {
	nodes []Node
}

func Run() {
	log.Println("DFS Struct")
	A := Node{node: "A", edges: []Node{}}
	B := Node{node: "B", edges: []Node{}}
	C := Node{node: "C", edges: []Node{}}
	D := Node{node: "D", edges: []Node{}}
	E := Node{node: "E", edges: []Node{}}
	F := Node{node: "F", edges: []Node{}}

	// Define the edges
	E.edges = []Node{F} // E connects to F
	C.edges = []Node{}
	B.edges = []Node{D, E}
	A.edges = []Node{B, C}

	// Create the graph
	graph = Graph{
		nodes: []Node{A, B, C, D, E, F},
	}
	DFS(A)
}

func DFS(start Node) {
	node := start
	visited[node.node] = true
	println("start, node, ", node.node)
	for _, edge := range node.edges {

		if !visited[edge.node] {
			DFS(edge)
		}
	}
}

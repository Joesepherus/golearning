package bfs

import (
	"golearning/queue"
	"log"
)

var graph = map[string][]string{
	"A": {"B", "C"},
	"B": {"D", "E"},
	"C": {},
	"D": {},
	"E": {"F"},
	"F": {},
}

var q = &queue.Queue{}

func Run() {
	log.Println("BFS")

	BFS("A")
}

func BFS(start string) {
	node := start
	q.Enqueue(node)
	log.Println("START NODE adding to queue: ", node)

	for !q.IsEmpty() {
		nodeAny, _ := q.Dequeue() // Dequeue returns an `any`

		node = nodeAny.(string)
		println("traversing node ", node)

		for _, edge := range graph[node] {
			log.Println("adding to queue: ", edge)
			q.Enqueue(edge)
		}

	}
}

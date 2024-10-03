package dfs

import "log"

var visited = make(map[string]bool)

var graph = map[string][]string{
	"A": {"B", "C"},
	"B": {"D", "E"},
	"C": {"F"},
	"D": {},
	"E": {"F"},
	"F": {},
}

func Run() {
	log.Println("DFS")

	DFS("A")
}

func DFS(start string) {
	visited[start] = true
	println("start, node, ", start)
	for _, edge := range graph[start] {

		if !visited[edge] {
			DFS(edge)
		}
	}
}

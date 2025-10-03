package main

import "fmt"

func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			fmt.Print(node, " ")
			visited[node] = true
			queue = append(queue, graph[node]...)
		}
	}
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {6},
		6: {},
	}

	bfs(graph, 1) // 1 2 3 4 5 6
}

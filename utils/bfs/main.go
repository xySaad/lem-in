package bfs

import (
	"slices"

	"lem-in/utils/parser"
)

func Bfs(af *parser.AntFarm) [][]string {
	visitedRooms := make(map[string]bool)
	queue := [][]string{}
	paths := [][]string{}
	visitedRooms[af.StartRoom] = true
	// add the start room to the queue as visted
	queue = append(queue, []string{af.StartRoom})

	for len(queue) > 0 {
		// Dequeue the current path
		path := queue[0]
		queue = queue[1:]
		endOfPath := path[len(path)-1]
		if endOfPath == af.EndRoom {
			paths = append(paths, path)
			continue
		}
		for room := range af.Rooms[endOfPath].Links {
			if !isVisited(path, room) && room != af.StartRoom {
				visitedRooms[room] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, room)
				queue = append(queue, newPath)
			}
			// fmt.Println(queue)
		}
	}

	return paths
}

func isVisited(path []string, room string) bool {
	return slices.Contains(path, room)
}

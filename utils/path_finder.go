package utils

import (
	"lem-in/utils/parser"
)

func FindAllPaths(af *parser.AntFarm) [][]string {
	paths := [][]string{{af.StartRoom}}
	queue := []string{af.StartRoom}
	visitedRooms := map[string]bool{}
	visitedRooms[af.StartRoom] = true
	conflictedPaths := [][]string{}
	for len(queue) > 0 {
		crawl(af, &paths, &queue, visitedRooms, &conflictedPaths)
	}

	return paths
}

func crawl(af *parser.AntFarm, paths *[][]string, queue *[]string, visitedRooms map[string]bool, conflictedPaths *[][]string) {
	foundaway := false
	for room := range af.Rooms[(*queue)[0]].Links {
		if visitedRooms[room] && room != af.EndRoom {
			continue
		}
		foundaway = true
		visitedRooms[room] = true
		newPath := append([]string{}, (*paths)[0]...) // Create a copy of the first path
		newPath = append(newPath, room)               // Append the new room
		*paths = append(*paths, newPath)
		*queue = append(*queue, room)
	}

	*queue = (*queue)[1:]

	if (*paths)[0][len((*paths)[0])-1] == af.EndRoom {
		*paths = append((*paths)[1:], (*paths)[0])
		return
	}
	if !foundaway {
		*conflictedPaths = append(*conflictedPaths, (*paths)[0])
	}
	*paths = (*paths)[1:]
}

package utils

import (
	"fmt"

	"lem-in/utils/parser"
)

func FindAllPaths(af *parser.AntFarm) map[string][][]string {
	paths := map[string][][]string{}
	for parrent := range af.Rooms[af.StartRoom].Links {
		paths[parrent] = nil
	}

	for parrent := range paths {
		for next := range af.Rooms[parrent].Links {
			if next == af.EndRoom || next == af.StartRoom {
				break
			}
			paths[parrent] = append(paths[parrent], []string{next})
			crawl(af, paths, parrent, next)
		}
	}
	return paths
}

func crawl(af *parser.AntFarm, paths map[string][][]string, current, next string) {
	for room := range af.Rooms[next].Links {
		_, ok := paths[room]
		if ok || room == current || room == af.StartRoom {
			continue
		}
		paths[current][len(paths[current])-1] = append(paths[current][len(paths[current])-1], room)
		fmt.Println(paths, current, next, room)
		if room == af.EndRoom {
			continue
		}
		crawl(af, paths, current, room)
	}
}

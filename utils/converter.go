package utils

import (
	"lem-in/utils/parser"
	"sort"
)

func ConvertPaths(af *parser.AntFarm, paths map[string][][]string) (ways [][]string) {
	DebugPrintf("paths: %v\n", paths)
	// Create a map to hold the smallest path for each starting link
	smallestPaths := make(map[string][]string)
	for startLink, way := range paths {
		if startLink == af.EndRoom {
			ways = append(ways, []string{af.StartRoom, startLink})
			continue
		}
		for _, path := range way {
			if len(path) > 0 && path[len(path)-1] == af.EndRoom {
				// Create the full path from startLink to the end
				fullPath := append([]string{af.StartRoom, startLink}, path...)
				// Check if we have a smaller path for this startLink
				if existingPath, exists := smallestPaths[startLink]; !exists || len(fullPath) < len(existingPath) {
					smallestPaths[startLink] = fullPath
				}
			}
		}
	}

	keys := []string{}
	for key := range smallestPaths {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// Collect the smallest paths into the ways slice
	for _, key := range keys {
		smallestPath := smallestPaths[key]
		ways = append(ways, smallestPath)
	}
	return
}

package utils

import (
	"fmt"
	"strconv"

	"lem-in/utils/parser"
)

func initializeAntMap(ways [][]string) map[int]int {
	antMap := make(map[int]int)
	for index := range ways {
		antMap[index] = len(ways[index])
	}
	return antMap
}

func findMinIndex(antMap map[int]int) int {
	minIndex := 0
	minValue := antMap[0]
	for index, value := range antMap {
		if value < minValue {
			minValue = value
			minIndex = index
		}
	}
	return minIndex
}

func initializeAntGroups(count int) [][]string {
	antGroups := make([][]string, count)
	for i := range antGroups {
		antGroups[i] = []string{}
	}
	return antGroups
}

func DistributeAnts(af *parser.AntFarm, paths map[string][][]string) {
	var ways [][]string
	debugPrintf("paths: %v\n", paths)
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

	// Collect the smallest paths into the ways slice
	for _, smallestPath := range smallestPaths {
		ways = append(ways, smallestPath)
	}

	debugPrintf("ways: %v\n", ways)

	antGroups := initializeAntGroups(len(ways))
	antID := 1
	antMap := initializeAntMap(ways)

	for antID <= af.Number {
		minIndex := findMinIndex(antMap)
		antGroups[minIndex] = append(antGroups[minIndex], "L"+strconv.Itoa(antID))
		antID++
		antMap[minIndex]++
	}
	manageTraffic(af, antGroups, ways)
}

func manageTraffic(af *parser.AntFarm, antGroups, ways [][]string) {
	traffic := make(map[string]int)
	unavailableRooms := make(map[string]bool)
	finishedAnts := []string{}
	steps := 0
	for len(finishedAnts) != af.Number {
		for i := 0; i < len(ways); i++ {
			unavailableRooms[af.EndRoom] = false
			for s := 0; s < len(antGroups[i]); s++ {
				ant := antGroups[i][s]
				if traffic[ant]+1 < len(ways[i]) && !unavailableRooms[ways[i][traffic[ant]+1]] {
					if ways[i][traffic[ant]+1] == af.EndRoom {
						unavailableRooms[ways[i][traffic[ant]]] = false
						finishedAnts = append(finishedAnts, ant)
						delete(traffic, ant)
						antGroups[i] = append(antGroups[i][:s], antGroups[i][s+1:]...)
						fmt.Printf("%v-%v ", ant, af.EndRoom)
						s--
						unavailableRooms[af.EndRoom] = true
					} else {
						fmt.Printf("%v-%v ", ant, ways[i][traffic[ant]+1])
						unavailableRooms[ways[i][traffic[ant]+1]] = true
						unavailableRooms[ways[i][traffic[ant]]] = false
						traffic[ant]++
					}
				}
			}
		}

		fmt.Println()
		steps++
	}
	debugPrint(steps)
}

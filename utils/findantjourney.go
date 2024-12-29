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

func DistributeAnts(af *parser.AntFarm, ways [][]string) {

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

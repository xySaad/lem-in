package pathfinder

import "lem-in/utils/parser"

type queued struct {
	room, parent string
}

type visitedRoom struct {
	parrent            string
	duplication, index int
}

type trackedRoom struct {
	name  string
	index int
}

type trackMap map[string][]trackedRoom
type Tunnels map[string][][]string

type PathFinder struct {
	AntFarm        *parser.AntFarm
	Tunnels        Tunnels
	Track          trackMap
	Queue          []queued
	Visited        map[string]*visitedRoom
	CurrentInQueue queued
	OptimalRoom    string
}

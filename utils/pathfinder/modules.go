package pathfinder

import (
	"lem-in/utils"
	"lem-in/utils/parser"
)

type queued struct {
	room, parent string
}

type visitedRoom struct {
	parrent            string
	duplication, index int
}

type (
	Tunnels map[string][]utils.Path
)

type PathFinder struct {
	AntFarm        *parser.AntFarm
	Tunnels        Tunnels
	Queue          []queued
	Visited        map[string]*visitedRoom
	CurrentInQueue queued
	OptimalRoom    string
}

func (p *PathFinder) ShiftTrack() {
	if len(p.CurrentPath().Track) == 0 {
		return
	}
	p.CurrentPath().Track = p.CurrentPath().Track[:len(p.CurrentPath().Track)-1]
}

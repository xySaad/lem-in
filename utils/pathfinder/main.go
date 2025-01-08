package pathfinder

import (
	"lem-in/utils"
	"lem-in/utils/parser"
)

func New(af *parser.AntFarm) (pf *PathFinder) {
	return &PathFinder{
		AntFarm: af,
		Tunnels: Tunnels{},
		Track:   trackMap{},
		Visited: map[string]*visitedRoom{},
	}
}

func (pf *PathFinder) init() {
	pf.Visited[pf.AntFarm.StartRoom] = &visitedRoom{}
	// Start from the starting room
	for link := range pf.AntFarm.Rooms[pf.AntFarm.StartRoom].Links {
		pf.Tunnels[link] = [][]string{}
		pf.Queue = append(pf.Queue, queued{parent: link, room: link})
		pf.Visited[link] = &visitedRoom{
			parrent: link,
		}
	}
}

func (pf *PathFinder) FindPaths() map[string][][]string {
	pf.init()
	for len(pf.Queue) > 0 {
		pf.CurrentInQueue = pf.Queue[0]
		pf.Queue = pf.Queue[1:]

		if len(pf.CurrentTunnel()) > 0 {
			if pf.LastRoom() != pf.CurrentInQueue.room {
				pf.Track.Shift(pf.CurrentInQueue.parent)
				continue
			} else if len(pf.CurrentPath()) == 0 {
				// if the current queued room has been removed (when a conflict found), delete it's path and skip it
				pf.Tunnels.Pop(pf.CurrentInQueue.parent)
				continue
			} else if len(pf.CurrentPath()) > 0 && pf.LastRoom() == pf.AntFarm.EndRoom {
				// if the current queued room is the end room append it to the end of the slice
				// "continue" to avoid ranging
				pf.Tunnels.RotateLeft(pf.CurrentInQueue.parent)
				utils.DebugPrint("found endroom => continue")
				continue
			}
		}

		utils.DebugPrint("current room:", pf.CurrentInQueue.room)
		utils.DebugPrint("initiail:", pf.Tunnels)

		// if the current room is the end skip (case: start is linked with the end directly)
		if pf.CurrentInQueue.room == pf.AntFarm.EndRoom {
			continue
		}
		// store paths to pass from when all links are conflicted
		optimalRoom := ""

		foundAway := pf.ForkPath()

		if foundAway {
			_, ok := pf.Visited[pf.CurrentInQueue.room]
			if !ok {
				panic(pf.CurrentInQueue.room)
			}
			pf.Visited[pf.CurrentInQueue.room].duplication--
		}

		utils.DebugPrint("[", foundAway, "]", "after iterating over links:")
		utils.DebugPrintf("paths: %v\n", pf.Tunnels)
		if !foundAway && len(pf.CurrentTunnel()) <= 1 {
			utils.DebugPrintf("optimalRoom: %v\n", optimalRoom)
			utils.DebugPrint(pf.CurrentInQueue.room, "didn't found a way")
			done := false
			roomToRemove, ok := pf.Visited[optimalRoom]
			if ok {
				utils.DebugPrintf("roomToRemove.duplication: %v\n", roomToRemove.duplication)
				utils.DebugPrint(roomToRemove.parrent != pf.CurrentInQueue.parent, len(pf.Tunnels[roomToRemove.parrent]) > roomToRemove.duplication)
			}
			if ok && roomToRemove.parrent != pf.CurrentInQueue.parent && len(pf.Tunnels[roomToRemove.parrent]) > roomToRemove.duplication {
				newGroupPaths := [][]string{}

				for i, conflictedPath := range pf.Tunnels[roomToRemove.parrent] {
					if roomToRemove.index < len(conflictedPath) && conflictedPath[roomToRemove.index] == optimalRoom {
						utils.DebugPrint("removing:", conflictedPath[roomToRemove.index])
						for _, r := range pf.Tunnels[roomToRemove.parrent][i] {
							if _, ok := pf.Visited[r]; ok && pf.Visited[r].duplication == 1 {
								utils.DebugPrint("remove visited room", r)
								delete(pf.Visited, r)
							} else if ok {
								utils.DebugPrintf("modify visited[r].duplication: %v --\n", r)
								pf.Visited[r].duplication--
								if r == "46" {
									utils.DebugPrintf("visited[r].duplication: %v\n", pf.Visited[r].duplication)
								}
							}
						}
						utils.DebugPrint("set to nil:", pf.Tunnels[roomToRemove.parrent][i])
						done = true
					} else {
						newGroupPaths = append(newGroupPaths, conflictedPath)
					}
				}
				pf.Tunnels[roomToRemove.parrent] = newGroupPaths
				if done {
					utils.DebugPrint("remove visited room", optimalRoom, "[optimal]")
					delete(pf.Visited, optimalRoom)
				}
			}
			if done {
				pf.Queue = append([]queued{pf.CurrentInQueue}, pf.Queue...)
				continue
			} else {
				// this is a dead end
				// TODO: back to the last room that has multiple links that are visited by other parrents,
				// if current parrent has only this path then find an exit (visited room)
				if len(pf.CurrentTunnel()) > 0 {
					utils.DebugPrint("dead end ", pf.CurrentInQueue.room, pf.CurrentPath(), pf.Track)
					utils.DebugPrintf("paths[current.parent][0]: %v\n", pf.CurrentPath())
					utils.DebugPrintf("track[current.parent]: %v\n", pf.ParrentTrack())
					nextRoomInQueue := pf.ParrentTrack()[len(pf.ParrentTrack())-1]

					if nextRoomInQueue.name == pf.CurrentInQueue.room {
						if len(pf.ParrentTrack()) <= 1 {
							utils.DebugPrint("optimal room:", optimalRoom)
							utils.DebugPrint("queue:")
							utils.DebugPrint("can't backtrack ", pf.ParrentTrack())
							continue
						}
						nextRoomInQueue = pf.ParrentTrack()[len(pf.ParrentTrack())-2]
					}

					pf.Queue = append([]queued{{
						room:   nextRoomInQueue.name,
						parent: pf.CurrentInQueue.parent,
					}}, pf.Queue...)

					pf.CurrentTunnel()[0] = pf.CurrentPath()[:nextRoomInQueue.index+1]
					utils.DebugPrintf("paths[current.parent][0]: %v\n", pf.CurrentPath())
					pf.Track[pf.CurrentInQueue.parent] = pf.ParrentTrack()[:len(pf.ParrentTrack())-1]
				}
			}
			if len(pf.CurrentTunnel()) > 0 {
				pf.Tunnels.RotateLeft(pf.CurrentInQueue.parent)
			} else {
				// if the room linked with the start room has dead end
				delete(pf.Tunnels, pf.CurrentInQueue.parent)
			}
			continue
		}

		if len(pf.CurrentTunnel()) > 0 && pf.CurrentInQueue.room != pf.CurrentInQueue.parent {
			if !foundAway {
				utils.DebugPrint("remove from visited beause it didn't split:", pf.LastRoom())
				for _, r := range pf.CurrentPath() {
					if pf.Visited[r].duplication == 1 {
						utils.DebugPrint("remove visited room", r)
						delete(pf.Visited, r)
					} else {
						utils.DebugPrintf("modify visited[r].duplication: %v --\n", r)
						pf.Visited[r].duplication--
						if r == "46" {
							utils.DebugPrintf("visited[r].duplication: %v\n", pf.Visited[r].duplication)
						}
					}
				}
			}
			pf.Tunnels.Pop(pf.CurrentInQueue.parent)
		}
		utils.DebugPrint("final:", pf.Tunnels)
	}
	return pf.Tunnels
}

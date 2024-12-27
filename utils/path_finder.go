package utils

import (
	"lem-in/utils/parser"
)

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

func FindPaths(af *parser.AntFarm) map[string][][]string {
	paths := make(map[string][][]string)
	track := make(map[string][]trackedRoom)
	queue := []queued{}
	visited := make(map[string]*visitedRoom)
	visited[af.StartRoom] = &visitedRoom{}
	// Start from the starting room
	for link := range af.Rooms[af.StartRoom].Links {
		paths[link] = [][]string{}
		queue = append(queue, queued{parent: link, room: link})
		visited[link] = &visitedRoom{
			parrent: link,
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		foundAway := false
		debugPrint("current room:", current.room)
		debugPrint("initiail:", paths)
		// if the current queued room has been removed (when a conflict found), delete it's path and skip it
		if len(paths[current.parent]) > 0 && len(paths[current.parent][0]) == 0 {
			debugPrint("removed empty path:", paths[current.parent][0])
			paths[current.parent] = paths[current.parent][1:]
			continue
		}

		// if the current queued room is the end room append it to the end of the slice
		// "continue" to avoid ranging
		if len(paths[current.parent]) > 0 && len(paths[current.parent][0]) > 0 && paths[current.parent][0][len(paths[current.parent][0])-1] == af.EndRoom {
			paths[current.parent] = append(paths[current.parent][1:], paths[current.parent][0])
			debugPrint("found endroom => continue")
			continue
		}
		// if the current room is the end skip (case: start is linked with the end directly)
		if current.room == af.EndRoom {
			continue
		}
		// store paths to pass from when all links are conflicted
		optimalRoom := ""

		// range over the links of current queued room
		for room := range af.Rooms[current.room].Links {
			vR, ok := visited[room]
			if ok && vR.parrent != current.parent && vR.parrent != "" && vR.parrent != af.StartRoom && len(paths[current.parent]) > 0 {
				debugPrint("possible way in parrent:", current.parent, "from:", current.room, "to:", room, "index:", vR.index, "\nroom visited in:", vR.parrent)
				if len(track[current.parent]) == 0 || track[current.parent][len(track[current.parent])-1].name != current.room {
					track[current.parent] = append(track[current.parent], trackedRoom{name: current.room, index: len(paths[current.parent][0]) - 1})
				}
			}
			if ok && room != af.EndRoom {
				_, ok := visited[optimalRoom]
				// biggest index is the best path to pass from
				if room != af.StartRoom && current.parent != vR.parrent && (!ok || vR.index > visited[optimalRoom].index) {
					debugPrintf("add optimal room: %v\n", room)
					debugPrintf("current.parent: %v\n", current.parent)
					debugPrintf("vR.parrent: %v\n", vR.parrent)
					optimalRoom = room
				}
				//TODO: handle the case where multiple rooms has the same index

				debugPrint("skipping room:", room, "link of:", current.room)
				continue
			}
			// mark the room as visited
			visited[room] = &visitedRoom{
				parrent:     current.parent,
				duplication: 1,
			}

			// the room is gonna be splitted as much as it's links
			// we increment duplication by one because a subpath is gonna be created
			_, ok = visited[current.room]
			if ok {
				visited[current.room].duplication++
			}

			foundAway = true
			if room == "a" {
				debugPrint("a found it's way through", room)
			}
			newPath := []string{}
			if len(paths[current.parent]) > 0 && current.room != current.parent {
				newPath = append(newPath, paths[current.parent][0]...)
			}
			if current.room == "a" {
				debugPrint("path to append:", newPath)
			}

			newPath = append(newPath, room)
			// store the index of the added room
			visited[room].index = len(newPath) - 1
			// append it
			paths[current.parent] = append(paths[current.parent], newPath)
			queue = append(queue, queued{parent: current.parent, room: room})
		}

		debugPrint("[", foundAway, "]", "after iterating over links:")
		debugPrintf("paths: %v\n", paths)
		if !foundAway && len(paths[current.parent]) <= 1 {
			debugPrintf("optimalRoom: %v\n", optimalRoom)
			debugPrint(current.room, "didn't found a way")
			done := false
			roomToRemove := visited[optimalRoom]
			debugPrint(roomToRemove.parrent != current.parent, len(paths[roomToRemove.parrent]) > roomToRemove.duplication)
			if roomToRemove.parrent != current.parent && len(paths[roomToRemove.parrent]) > roomToRemove.duplication {
				for i, conflictedPath := range paths[roomToRemove.parrent] {
					if roomToRemove.index < len(conflictedPath) && conflictedPath[roomToRemove.index] == optimalRoom {
						debugPrint("removing:", conflictedPath[roomToRemove.index])
						for _, r := range paths[roomToRemove.parrent][i] {
							delete(visited, r)
						}
						paths[roomToRemove.parrent][i] = nil
						done = true
					}
				}
				if done {
					delete(visited, optimalRoom)
				}
			}
			if done {
				queue = append([]queued{current}, queue...)
				continue
			} else {
				// this is a dead end
				// TODO: back to the last room that has multiple links that are visited by other parrents,
				// if current parrent has only this path then find an exit (visited room)
				if len(paths[current.parent]) > 0 {
					debugPrint("dead end ", current.room, paths[current.parent][0], track)
					debugPrintf("paths[current.parent][0]: %v\n", paths[current.parent][0])
					debugPrintf("track[current.parent]: %v\n", track[current.parent])
					nextRoomInQueue := track[current.parent][len(track[current.parent])-1]

					if nextRoomInQueue.name == current.room {
						if len(track[current.parent]) <= 1 {
							debugPrint("can't backtrack ", track[current.parent])
							continue
						}
						nextRoomInQueue = track[current.parent][len(track[current.parent])-2]
					}

					queue = append([]queued{{
						room:   nextRoomInQueue.name,
						parent: current.parent,
					}}, queue...)

					paths[current.parent][0] = paths[current.parent][0][:nextRoomInQueue.index+1]
					debugPrintf("paths[current.parent][0]: %v\n", paths[current.parent][0])
					track[current.parent] = track[current.parent][:len(track[current.parent])-1]
				}
			}
			if len(paths[current.parent]) > 0 {
				paths[current.parent] = append(paths[current.parent][1:], paths[current.parent][0])
			} else {
				// if the room linked with the start room has dead end
				delete(paths, current.parent)
			}
			continue
		}
		// the path foundaway means it got splited into subpath
		// it's parrent path has been removed so we decrement one
		_, ok := visited[current.room]
		if ok {
			visited[current.room].duplication--
		}
		if len(paths[current.parent]) > 0 && current.room != current.parent {
			paths[current.parent] = paths[current.parent][1:]
		}
		debugPrint("final:", paths)
	}

	return paths
}

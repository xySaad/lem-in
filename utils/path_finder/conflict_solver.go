package path_finder

func (li *LemIn) Escape() bool {
	currentParrent := li.current.parrent
	currentRoom := li.current.room
	optimalRoom := li.optimalRoom
	if !li.foundAway && li.Tunnels.Len(currentParrent) <= 1 {
		// debugPrintf("optimalRoom: %v\n", li.optimalRoom)
		// debugPrint(currentRoom, "didn't found a way")
		done := false
		roomToRemove, ok := li.visited[optimalRoom]
		if ok {
			// debugPrintf("roomToRemove.duplication: %v\n", roomToRemove.duplication)
			// debugPrint(roomToRemove.parrent != currentParrent, len(paths[roomToRemove.parrent]) > roomToRemove.duplication)
		}
		if ok && roomToRemove.parrent != currentParrent && li.Tunnels.Len(roomToRemove.parrent) > roomToRemove.duplication {
			// newGroupPaths := [][]string{}

			for i, conflictedPath := range *li.Tunnels[roomToRemove.parrent].All() {
				if roomToRemove.index < conflictedPath.Len() && *conflictedPath.removeAt(roomToRemove.index) == optimalRoom {
					delete(li.visited, optimalRoom)
					conflictedPath = *li.Tunnels[roomToRemove.parrent].removeAt(i)
					// debugPrint("removing:", conflictedPath[roomToRemove.index])
					for _, r := range *conflictedPath.All() {
						if _, ok := li.visited[r]; ok && li.visited[r].duplication == 1 {
							// debugPrint("remove visited room", r)
							delete(li.visited, r)
						} else if ok {
							// debugPrintf("modify li.visited[r].duplication: %v --\n", r)
							li.visited[r].duplication--
						}
					}
					// debugPrint("set to nil:", paths[roomToRemove.parrent][i])
					done = true
				} else {
					// newGroupPaths = append(newGroupPaths, conflictedPath)
				}
			}
			// paths[roomToRemove.parrent] = newGroupPaths
			if done {
				// debugPrint("remove visited room", optimalRoom, "[optimal]")
				delete(li.visited, optimalRoom)
			}
		}
		if done {
			li.queue.Append(li.current)
			return true
		} else {
			// this is a dead end
			// TODO: back to the last room that has multiple links that are visited by other parrents,
			// if current parrent has only this path then find an exit (visited room)
			if li.Tunnels[currentParrent].Len() > 0 {
				// debugPrint("dead end ", currentRoom, paths[currentParrent][0], track)
				// debugPrintf("paths[currentParrent][0]: %v\n", paths[currentParrent][0])
				// debugPrintf("track[currentParrent]: %v\n", track[currentParrent])
				nextRoomInQueue, _ := li.track[currentParrent].Pop()

				if nextRoomInQueue.name == currentRoom {
					if li.track[currentParrent].Len() <= 1 {
						// debugPrint("optimal room:", optimalRoom)
						// debugPrint("queue:")
						// debugPrint("can't backtrack ", track[currentParrent])
						return true
					}
					nextRoomInQueue, _ = li.track[currentParrent].Pop()
				}

				li.queue.Append(queued{
					room:    nextRoomInQueue.name,
					parrent: currentParrent,
				})
				bottom := li.Tunnels.Bottom(currentParrent)
				bottom.slice = bottom.slice[:nextRoomInQueue.index+1]
				// paths[currentParrent][0] = paths[currentParrent][0][:nextRoomInQueue.index+1]
				// debugPrintf("paths[currentParrent][0]: %v\n", paths[currentParrent][0])
			}
		}
		if li.Tunnels.Len(currentParrent) > 0 {
			li.Tunnels[currentParrent].Push(*li.Tunnels[currentParrent].Shift())
		} else {
			// if the room linked with the start room has dead end
			delete(li.Tunnels, currentParrent)
		}
		return true
	}
	return true
}

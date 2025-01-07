package path_finder

func (li *LemIn) run() (doContinue bool) {

	doContinue = true
	if li.queue.Len() <= 0 {
		return false
	}
	li.current = *li.queue.Shift()
	currentRoom := li.current.room
	currentParrent := li.current.parrent
	currentTunnel := li.Tunnels[currentParrent]
	currentTrack := li.track[currentParrent]
	bottom := li.Tunnels.Bottom(currentParrent)

	if currentTunnel.Len() > 0 && *bottom.Peek() != currentRoom {
		currentTrack.Pop()
		return
	}

	if currentTunnel.Len() > 0 && bottom.Len() == 0 {
		currentTunnel.Pop()
		return
	}

	// if the current queued room is the end room append it to the end of the slice
	// "continue" to avoid ranging
	if bottom.Peek() != nil && *bottom.Peek() == li.af.EndRoom {
		currentTunnel.Push(*currentTunnel.Shift())
		return
	}

	return li.crawl()
}

func (li *LemIn) crawl() bool {
	// range over the links of current queued room
	currentParrent := li.current.parrent
	currentTunnel := li.Tunnels[currentParrent]

	currentRoom := li.current.room
	bottom := li.Tunnels.Bottom(currentParrent)

	for room := range li.af.Rooms[currentRoom].Links {
		vR, ok := li.visited[room]
		currentTrack := li.track[currentParrent]

		if ok && vR.parrent != currentParrent && vR.parrent != "" && vR.parrent != li.af.StartRoom {
			// debugPrint("possible way in parrent:", currentParrent, "from:", currentRoom, "to:", room, "index:", vR.index, "\nroom visited in:", vR.parrent)
			topInTrack, exists := currentTrack.Pop()
			if !exists || topInTrack.name != currentRoom {
				currentTrack.Push(trackedRoom{name: currentRoom, index: bottom.Len() - 1})
			}
		}
		if ok && room != li.af.EndRoom {
			_, ok := li.visited[li.optimalRoom]
			// biggest index is the best path to pass from
			if room != li.af.StartRoom && currentParrent != vR.parrent && (!ok || vR.index > li.visited[li.optimalRoom].index) {
				// debugPrintf("add optimal room: %v\n", room)
				// debugPrintf("currentParrent: %v\n", currentParrent)
				// debugPrintf("vR.parrent: %v\n", vR.parrent)
				li.optimalRoom = room
			}
			//TODO: handle the case where multiple rooms has the same index

			// debugPrint("skipping room:", room, "link of:", currentRoom)
			continue
		}
		// mark the room as visited
		li.visited[room] = &visitedRoom{
			parrent:     currentParrent,
			duplication: 0,
		}

		li.foundAway = true

		newPath := dequeue[string]{}
		if currentRoom != currentParrent {
			newPath = *bottom
		}

		newPath.Push(room)
		for _, r := range *newPath.All() {
			// debugPrintf("modify visited[r].duplication: %v ++\n", r)
			li.visited[r].duplication++
		}
		// store the index of the added room
		li.visited[room].index = newPath.Len() - 1
		// append it
		currentTunnel.Push(newPath)
		li.queue.Push(queued{parrent: currentParrent, room: room})
	}
	if li.foundAway {
		li.visited[currentRoom].duplication--
	} else if currentTunnel.Len() <= 1 {
		return li.Escape()
	}
	return true
}

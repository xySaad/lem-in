package pathfinder

import (
	"lem-in/utils"
)

func (pf *PathFinder) IsOptimal(link string) bool {
	// biggest index is the best path to pass from

	visitedLink := pf.Visited[link]
	optimalRoom, ok := pf.Visited[pf.OptimalRoom]

	return link != pf.AntFarm.StartRoom &&
		pf.CurrentInQueue.parent != visitedLink.parrent &&
		(!ok || visitedLink.index > optimalRoom.index)
}

// range over the links of current queued room
func (pf *PathFinder) ForkPath() (foundAway bool) {
	currentParrent := pf.CurrentInQueue.parent
	currentRoom := pf.CurrentInQueue.room

	for link := range pf.AntFarm.Rooms[currentRoom].Links {

		visitedLink, ok := pf.Visited[link]

		if ok {
			if visitedLink.parrent != currentParrent &&
				visitedLink.parrent != "" &&
				visitedLink.parrent != pf.AntFarm.StartRoom &&
				len(pf.CurrentTunnel()) > 0 &&
				pf.LastTrack().name != currentRoom {
				utils.DebugPrint("possible way in parrent:", currentParrent, "from:", currentRoom, "to:", link, "index:", visitedLink.index, "\nroom visited in:", visitedLink.parrent)
				pf.Track[currentParrent] = append(pf.ParrentTrack(), trackedRoom{name: currentRoom, index: len(pf.CurrentPath()) - 1})
			}

			if link != pf.AntFarm.EndRoom {
				if pf.IsOptimal(link) {
					utils.DebugPrintf("add optimal room: %v\n", link)
					utils.DebugPrintf("current.parent: %v\n", currentParrent)
					utils.DebugPrintf("vR.parrent: %v\n", visitedLink.parrent)
					pf.OptimalRoom = link
				}
				//TODO: handle the case where multiple rooms has the same index

				utils.DebugPrint("skipping room:", link, "link of:", currentRoom)
				continue
			}
		}

		// mark the room as visited
		pf.Visited[link] = &visitedRoom{
			parrent:     currentParrent,
			duplication: 0,
		}

		foundAway = true

		newPath := []string{}
		if len(pf.CurrentTunnel()) > 0 && currentRoom != currentParrent {
			newPath = append(newPath, pf.CurrentPath()...)
		}

		newPath = append(newPath, link)
		for _, r := range newPath {
			utils.DebugPrintf("modify visited[r].duplication: %v ++\n", r)
			pf.Visited[r].duplication++
			if r == "46" {
				utils.DebugPrintf("visited[r].duplication: %v\n", pf.Visited[r].duplication)
			}
		}
		// store the index of the added room
		pf.Visited[link].index = len(newPath) - 1
		// append it
		pf.Tunnels[currentParrent] = append(pf.CurrentTunnel(), newPath)
		pf.Queue = append(pf.Queue, queued{parent: currentParrent, room: link})
	}
	return
}
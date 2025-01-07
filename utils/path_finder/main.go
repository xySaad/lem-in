package path_finder

import (
	"lem-in/utils/parser"
)

func (li *LemIn) init(af *parser.AntFarm) {
	// Initial structure
	li.Tunnels = Tunnels{}
	li.af = af
	li.track = map[string]*dequeue[trackedRoom]{}
	li.visited = map[string]*visitedRoom{}
	li.queue = *NewQueue([]queued{})

	// Mark start room as visited
	li.visited[af.StartRoom] = &visitedRoom{}
	// Visit links of starting room
	for link := range af.Rooms[af.StartRoom].Links {
		// Create a tunnel for each link
		li.Tunnels[link] = NewTunnel()
		li.queue.Push(queued{parrent: link, room: link})
		li.visited[link] = &visitedRoom{parrent: link}
	}

}

func FindPaths(af *parser.AntFarm) map[string][][]string {
	lemin := LemIn{}
	lemin.init(af)

	for lemin.run() {
		if lemin.Len(lemin.current.parrent) > 0 && lemin.current.room != lemin.current.parrent {
			if !lemin.foundAway {
				// debugPrint("remove from visited beause it didn't split:", paths[current.parent][0][len(paths[current.parent][0])-1])
				bottom := lemin.Tunnels[lemin.current.parrent].Shift()
				for _, r := range *bottom.All() {
					if lemin.visited[r].duplication == 1 {
						// debugPrint("remove visited room", r)
						delete(lemin.visited, r)
					} else {
						// debugPrintf("modify visited[r].duplication: %v --\n", r)
						lemin.visited[r].duplication--
					}
				}
			}
		}
	}
	result := map[string][][]string{}
	for parrent, tunnel := range lemin.Tunnels {
		for _, pathDq := range *tunnel.All() {
			result[parrent] = append(result[parrent], pathDq.slice)
		}
	}
	return result
}

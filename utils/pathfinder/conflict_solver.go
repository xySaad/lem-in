package pathfinder

func (pf *PathFinder) RemoveConflictedPath(conflictedParrent string, indexOfPath int) {
	for _, r := range pf.Tunnels[conflictedParrent][indexOfPath].Route {
		if pf.Visited[r].duplication <= 1 {
			// utils.DebugPrint("remove visited room", r)
			delete(pf.Visited, r)
		} else {
			// utils.DebugPrintf("modify visited[r].duplication: %v --\n", r)
			pf.Visited[r].duplication--
		}
	}
}

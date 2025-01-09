package pathfinder

func (pf *PathFinder) CurrentTunnel() [][]string {
	return pf.Tunnels[pf.CurrentParrent()]
}

func (pf *PathFinder) CurrentPath() []string {
	return pf.CurrentTunnel()[0]
}
func (pf *PathFinder) LastRoom() string {
	return pf.CurrentPath()[len(pf.CurrentPath())-1]
}

func (pf *PathFinder) ParrentTrack() []trackedRoom {
	return pf.Track[pf.CurrentParrent()]
}

func (pf *PathFinder) LastTrack() trackedRoom {
	if len(pf.ParrentTrack()) == 0 {
		return trackedRoom{}
	}
	return pf.ParrentTrack()[len(pf.ParrentTrack())-1]
}

func (pf *PathFinder) CurrentParrent() string {
	return pf.CurrentInQueue.parent
}

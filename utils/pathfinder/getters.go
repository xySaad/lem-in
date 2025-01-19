package pathfinder

import "lem-in/utils"

func (pf *PathFinder) CurrentTunnel() []utils.Path {
	return pf.Tunnels[pf.CurrentParrent()]
}

func (pf *PathFinder) CurrentPath() *utils.Path {
	return &pf.CurrentTunnel()[0]
}

func (pf *PathFinder) LastRoom() string {
	return pf.CurrentPath().Route[len(pf.CurrentPath().Route)-1]
}

func (pf *PathFinder) ParrentTrack() []utils.TrackedRoom {
	if len(pf.CurrentTunnel()) == 0 {
		return []utils.TrackedRoom{}
	}
	return pf.CurrentPath().Track
}

func (pf *PathFinder) LastTrack() utils.TrackedRoom {
	if len(pf.ParrentTrack()) == 0 {
		return utils.TrackedRoom{}
	}
	return pf.ParrentTrack()[len(pf.ParrentTrack())-1]
}

func (pf *PathFinder) CurrentParrent() string {
	return pf.CurrentInQueue.parent
}

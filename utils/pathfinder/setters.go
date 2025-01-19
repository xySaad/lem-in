package pathfinder

func (tnl *Tunnels) Pop(parrent string) {
	if len((*tnl)[parrent]) == 0 {
		return
	}
	(*tnl)[parrent] = (*tnl)[parrent][1:]
}

func (tnl *Tunnels) RotateLeft(parrent string) {
	(*tnl)[parrent] = append((*tnl)[parrent][1:], (*tnl)[parrent][0])
}

func (pf *PathFinder) AppendQueue(x queued) {
	pf.Queue = append(pf.Queue, x)
}

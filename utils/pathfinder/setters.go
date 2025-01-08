package pathfinder

func (tm *trackMap) Shift(parrent string) {
	if len((*tm)[parrent]) == 0 {
		return
	}
	(*tm)[parrent] = (*tm)[parrent][:len((*tm)[parrent])-1]
}

func (tnl *Tunnels) Pop(parrent string) {
	if len((*tnl)[parrent]) == 0 {
		return
	}
	(*tnl)[parrent] = (*tnl)[parrent][1:]
}

func (tnl *Tunnels) RotateLeft(parrent string) {
	(*tnl)[parrent] = append((*tnl)[parrent][1:], (*tnl)[parrent][0])
}

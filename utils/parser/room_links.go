package parser

func (af *antFarm) parseRoomLinks() error {
	for i, char := range af.currentLine {
		if char == '-' {
			room, exists := af.rooms[af.currentLine[:i]]
			if exists {
				_, exists := af.rooms[af.currentLine[i+1:]]
				if exists {
					room.links = append(room.links, af.currentLine[i+1:])
				}
			}
		}
	}
	return nil
}

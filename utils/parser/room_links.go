package parser

func (af *antFarm) parseRoomLinks() error {
	for i, char := range af.currentLine {
		if char == '-' {
			room, exists := af.rooms[af.currentLine[:i]]
			if !exists {
				return af.ParsingError("can't link unexisted room", 0)
			}
			_, exists = af.rooms[af.currentLine[i+1:]]
			if !exists {
				return af.ParsingError("can't link unexisted room", 0)
			}
			room.links = append(room.links, af.currentLine[i+1:])

		}
	}
	return nil
}

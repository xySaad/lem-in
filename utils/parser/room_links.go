package parser

func (af *antFarm) parseRoomLinks() error {
	for i, char := range af.currentLine {
		if char == '-' {
			if i == len(af.currentLine)-1 {
				return af.ParsingError("invalid link format", 0)
			}
			if af.currentLine[:i] == af.currentLine[i+1:] {
				return af.ParsingError("can't link a room with itself", 0)
			}
			room, exists := af.rooms[af.currentLine[:i]]
			if !exists {
				return af.ParsingError("can't link unexisted room", 0)
			}
			_, exists = af.rooms[af.currentLine[i+1:]]
			if !exists {
				return af.ParsingError("can't link unexisted room", 0)
			}
			_, ok := room.links[af.currentLine[i+1:]]
			if ok {
				return af.ParsingError("duplicated links", 0)
			}
			room.links[af.currentLine[i+1:]] = struct{}{}
			af.state.prevToken = dash
			break
		}
	}
	if af.state.prevToken != dash {
		return af.ParsingError("invalid link", 0)
	}
	return nil
}

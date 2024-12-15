package parser

func (af *antFarm) parseRoomList() error {
	spaceIndex := 0
	for i, char := range af.currentLine {
		switch af.state.expectedToken {
		case roomCharacter:
			if af.state.prevToken == space {
				return af.ParsingError("missing y coordinates", i)
			}
			if char == ' ' {
				spaceIndex = i
				_, exist := af.rooms[af.currentLine[:i]]
				if exist {
					return af.ParsingError("duplicated room", 0)
				}
				af.rooms[af.currentLine[:i]] = &room{
					links: map[string]struct{}{},
				}
				if af.state.prevState == start {
					af.startRoom = af.currentLine[:i]
				}
				if af.state.prevState == end {
					af.endRoom = af.currentLine[:i]
				}
				af.state.prevToken = space
				af.state.expectedToken = x
				continue
			}
			if char == '-' && i != 0 {
				if af.startRoom == "" || af.endRoom == "" {
					return af.ParsingError("no start/end room provided", 0)
				}
				if af.startRoom == "" {
					return af.ParsingError("no start room")
				}
				af.state.prevState = roomsList
				af.state.expectedState = roomLinks
				return nil
			}
		case x:
			if char == ' ' {
				if af.state.prevToken == space {
					return af.ParsingError("invalid space after", i)
				}
				af.state.prevToken = space
				af.state.expectedToken = y

				continue
			}
			if char >= '0' && char <= '9' {
				af.rooms[af.currentLine[:spaceIndex]].x = af.rooms[af.currentLine[:spaceIndex]].x*10 + int(char-'0')
				af.state.prevToken = x
			} else {
				return af.ParsingError("invalid x value", i)
			}
		case y:
			if char == ' ' {
				if af.state.prevToken == space || i < len(af.currentLine) {
					return af.ParsingError("invalid space after", i)
				}
				af.state.prevToken = space
				af.state.expectedToken = roomCharacter
				continue
			}
			if char >= '0' && char <= '9' {
				af.rooms[af.currentLine[:spaceIndex]].y = af.rooms[af.currentLine[:spaceIndex]].y*10 + int(char-'0')
				af.state.prevToken = y
			} else {
				return af.ParsingError("invalid y value", i)
			}
		}
	}
	if af.state.prevToken == y {
		room, alo := af.rooms[af.currentLine[:spaceIndex]]
		if !alo {
			return af.ParsingError("invalid format", 0)
		}
		if room.x < room.y {
			uniquePair := room.y*room.y + room.x
			_, exists := af.xyPairs[uniquePair]
			if exists {
				return af.ParsingError("duplicated coordinates", 0)
			}
			af.xyPairs[uniquePair] = struct{}{}
		} else {
			uniquePair := room.x*room.x + room.x + room.y
			_, exists := af.xyPairs[uniquePair]
			if exists {
				return af.ParsingError("duplicated coordinates", 0)
			}
			af.xyPairs[uniquePair] = struct{}{}
		}
	}
	if af.currentLine[:spaceIndex] == "" && af.state.expectedToken == roomCharacter {
		return af.ParsingError("invalid format", 0)
	}
	return nil
}

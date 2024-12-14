package parser

func (af *antFarm) parseRoomList() error {
	roomName := []rune{}
	roomStr := ""
	for i, char := range af.currentLine {
		switch af.state.expectedToken {
		case roomCharacter:
			if af.state.prevToken == space {
				return af.ParsingError("missing y coordinates", i)
			}
			if char == ' ' {
				roomStr = string(roomName)
				_, exist := af.rooms[roomStr]
				if exist {
					return af.ParsingError("duplicated room", 0)
				}
				af.rooms[roomStr] = &room{
					x: -1,
					y: -1,
				}
				if af.state.prevState == start {
					af.startRoom = roomStr
				}
				if af.state.prevState == end {
					af.endRoom = roomStr
				}
				af.state.prevToken = space
				af.state.expectedToken = x
				continue
			}
			if char == '-' {
				if af.startRoom == "" {
					return af.ParsingError("no start room")
				}
				af.state.prevState = roomsList
				af.state.expectedState = links
				return nil
			}
			roomName = append(roomName, char)
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
				af.rooms[roomStr].x = af.rooms[roomStr].x*10 + int(char-'0')
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
				af.rooms[roomStr].y = af.rooms[roomStr].y*10 + int(char-'0')
				af.state.prevToken = y
			} else {
				return af.ParsingError("invalid y value", i)
			}
		}
	}
	if af.state.prevToken == y {
		room, alo := af.rooms[roomStr]
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
	if roomStr == "" && af.state.expectedToken == roomCharacter {
		return af.ParsingError("invalid format", 0)
	}
	return nil
}

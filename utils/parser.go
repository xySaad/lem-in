package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	// states
	antsNumber = iota
	start
	roomsList
	end
	links
	// tokens
	roomCharacter
	x
	y
	space
	dash    = '-'
	hashtag = '#'
)

type state struct {
	prevToken, expectedToken, prevState, expectedState, linePosition int
}

type antFarm struct {
	xyPairs            map[int]struct{}
	rooms              map[string]*room
	startRoom, endRoom string
	state              *state
	currentLine        string
}

type room struct {
	x, y int
}

func initFarm() antFarm {
	return antFarm{
		xyPairs: map[int]struct{}{},
		rooms:   map[string]*room{},
		state: &state{
			expectedState: antsNumber,
		},
	}
}

func ParseFile(filename string) error {
	af := initFarm()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	buff := make([]byte, 1)
	line := []byte{}

	for {
		_, err := file.Read(buff)
		if err != nil && err != io.EOF {
			return err
		}

		if buff[0] == '\n' {
			af.currentLine = string(line)
			af.state.linePosition++

			if len(line) == 0 {
				return af.ParsingError("invalid empty line", 0)
			}
			if line[0] == '#' {
				if string(line) == "##start" {
					if af.startRoom != "" || af.state.prevState == start {
						return af.ParsingError("duplicated start room")
					}
					af.state.prevState = start
					af.state.expectedState = roomsList
					af.state.expectedToken = roomCharacter
				}
				if string(line) == "##end" {
					if af.endRoom != "" || af.state.prevState == end {
						return af.ParsingError("duplicated end room")
					}
					if af.state.prevState == start {
						return af.ParsingError("no start room provided")
					}
					af.state.expectedState = roomsList
					af.state.expectedToken = roomCharacter
					af.state.prevState = end
				}
				line = nil
				continue
			}
			err = af.parseLine()
			if err != nil {
				return err
			}
			line = nil
		} else {
			line = append(line, buff[0])
		}

		if err == io.EOF {
			break
		}
	}
	return nil
}

func (af *antFarm) parseLine() error {
	lineStr := af.currentLine
	switch af.state.expectedState {
	case antsNumber:
		n, err := strconv.Atoi(lineStr)
		if err != nil || n <= 0 {
			return errors.New("invalid ants number")
		}
		af.state.prevState = antsNumber
		af.state.expectedState = roomsList
	case roomsList:
		defer func() {
			af.state.prevState = roomsList
			af.state.expectedToken = roomCharacter
		}()
		return af.checkCoords()
	}
	return nil
}

func (af *antFarm) checkCoords() error {
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
				af.rooms[roomStr] = &room{}
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
				af.state.prevToken = dash
				af.state.prevState = roomsList
				af.state.expectedState = links
				continue
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
			fmt.Println(roomStr)
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
	return nil
}

func (af *antFarm) ParsingError(err string, i ...int) error {
	if len(i) == 1 {
		index := i[0]
		//TODO: show the word where the error is
		err += fmt.Sprint(": ", "\""+af.currentLine+"\" at ", af.state.linePosition, ":", index)
	}

	return errors.New(err)
}

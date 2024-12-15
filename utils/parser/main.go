package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

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
			if af.state.linePosition == 1 {
				n, err := strconv.Atoi(af.currentLine)
				if err != nil || n <= 0 {
					return errors.New("invalid ants number")
				}
				af.number = n
				af.state.prevToken = antsNumber
				af.state.expectedState = roomsList
			}
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
	fmt.Println("start:", af.startRoom, "end:", af.endRoom)

	for name, room := range af.rooms {
		fmt.Print("room:", name, " x:", room.x, " y:", room.y, " links:", room.links, "\n")
	}
	return nil
}

func (af *antFarm) parseLine() error {
	switch af.state.expectedState {
	case roomsList:
		defer func() {
			af.state.prevState = roomsList
			af.state.expectedToken = roomCharacter
		}()
		return af.parseRoomList()
	case roomLinks:
		af.state.prevToken = roomCharacter
		return af.parseRoomLinks()
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

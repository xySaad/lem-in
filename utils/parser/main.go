package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ParseFile(filename string) (*AntFarm, error) {
	af := initFarm()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buff := make([]byte, 1)
	line := []byte{}

	for {
		_, err := file.Read(buff)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if buff[0] == '\n' {
			if len(line) == 0 {
				continue
			}
			af.currentLine = string(line)
			af.state.linePosition++
			if af.state.linePosition == 1 {
				n, err := strconv.Atoi(af.currentLine)
				if err != nil || n <= 0 {
					return nil, errors.New("invalid ants number")
				}
				af.number = n
				af.state.prevToken = antsNumber
				af.state.expectedState = roomsList
			}
			if line[0] == '#' {
				if string(line) == "##start" {
					if af.StartRoom != "" || af.state.prevState == start {
						return nil, af.parsingError("duplicated start room")
					}
					af.state.prevState = start
					af.state.expectedState = roomsList
					af.state.expectedToken = roomCharacter
				}
				if string(line) == "##end" {
					if af.EndRoom != "" || af.state.prevState == end {
						return nil, af.parsingError("duplicated end room")
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
				return nil, err
			}
			line = nil
		} else {
			line = append(line, buff[0])
		}

		if err == io.EOF {
			break
		}
	}
	return af, nil
}

func (af *AntFarm) parseLine() error {
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

func (af *AntFarm) parsingError(err string, i ...int) error {
	if len(i) == 1 {
		index := i[0]
		// TODO: show the word where the error is
		err += fmt.Sprint(": ", "\""+af.currentLine+"\" at ", af.state.linePosition, ":", index)
	}

	return errors.New(err)
}

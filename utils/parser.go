package utils

import (
	"errors"
	"io"
	"os"
	"strconv"
)

const (
	// states
	antsNumber = 1 << iota
	start
	coords
	end
	links
	// tokens
	room
	space
	dash    = '-'
	hashtag = '#'
)

type state struct {
	prevToken, expectedToken int
	currentState             int
}

type antFarm struct {
	state        *state
	currentChunk []byte
}

func initFarm() antFarm {
	return antFarm{
		state: &state{
			currentState: antsNumber,
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
			err = af.parseLine(line)
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

func (af *antFarm) parseLine(line []byte) error {
	lineStr := string(line)
	switch af.state.currentState {
	case antsNumber:
		n, err := strconv.Atoi(lineStr)
		if n <= 0 {
			return errors.New("invalid ant number: " + lineStr)
		}
		if err != nil {
			return errors.New("invalid ant number: " + lineStr)
		}
		af.state.currentState = start
	case start:
		if lineStr == "##start" {
			af.state.currentState = coords
			af.state.expectedToken = room
		} else {
			return errors.New("expected ##start found " + lineStr)
		}
	case coords:
		if lineStr == "##end" {
			if af.state.prevToken == room {
				return errors.New("[lem-in] invalid format: no start room")
			}
			af.state.currentState = links
			return nil
		}
		var sections int
		for _, char := range lineStr {
			switch af.state.expectedToken {
			case room:
				if char != 'L' && char != '#' {
					af.state.expectedToken = room | space
					af.state.prevToken = room
				} else {
					return errors.New("expected coordinate found: " + lineStr)
				}
			case room | space:
				if char == ' ' && af.state.prevToken != space {
					sections++
					if sections > 2 {
						return errors.New("expected new line found: " + lineStr)
					}
					af.state.currentState = room
					af.state.prevToken = space
					continue
				} else if char >= '0' && char <= '9' {
					af.state.expectedToken = space | room
					af.state.prevToken = room
				} else if sections > 0 {
					return errors.New("expected coordinate found: " + lineStr)
				}
			}
		}
		af.state.currentState = coords
	}
	return nil
}

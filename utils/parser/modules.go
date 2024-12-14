package parser

const (
	// states
	start = iota
	roomsList
	end
	links
	// tokens
	antsNumber
	roomCharacter
	x
	y
	space
	dash
	hashtag = '#'
)

type state struct {
	prevToken, expectedToken, prevState, expectedState, linePosition int
}

type antFarm struct {
	number             int
	xyPairs            map[int]struct{}
	rooms              map[string]*room
	startRoom, endRoom string
	state              *state
	currentLine        string
}

type room struct {
	links []string
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

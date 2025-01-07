package path_finder

import (
	"lem-in/utils/parser"
)

type visitedRoom struct {
	parrent            string
	duplication, index int
}

type trackedRoom struct {
	name  string
	index int
}

type Tunnels map[string]*dequeue[dequeue[string]]

func NewTunnel() *dequeue[dequeue[string]] {
	return &dequeue[dequeue[string]]{}
}

func (p Tunnels) Len(parrent string) int {
	return p[parrent].Len()
}
func (p Tunnels) Peek(parrent string) *dequeue[string] {
	return p[parrent].Peek()
}
func (p Tunnels) Bottom(parrent string) *dequeue[string] {
	return p[parrent].Bottom()
}

type LemIn struct {
	Tunnels
	track       map[string]*dequeue[trackedRoom]
	queue       dequeue[queued]
	visited     map[string]*visitedRoom
	af          *parser.AntFarm
	current     queued
	foundAway   bool
	optimalRoom string
}

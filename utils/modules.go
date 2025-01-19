package utils

type TrackedRoom struct {
	Name  string
	Index int
}

type Path struct {
	Route []string
	Track []TrackedRoom
}

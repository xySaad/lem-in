package main

import (
	"fmt"
	"lem-in/utils"
	"lem-in/utils/parser"
	"lem-in/utils/pathfinder"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		fmt.Println("usage: go run . <filename.txt>")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("missing arguments")
		fmt.Println("usage: go run . <filename.txt>")
		return
	}
	antFarm, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	pf := pathfinder.New(antFarm)

	paths := pf.FindPaths()

	utils.DistributeAnts(antFarm, utils.ConvertPaths(antFarm, paths))
}

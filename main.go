package main

import (
	"fmt"
	"os"

	"lem-in/utils"
	"lem-in/utils/parser"
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

	paths := utils.FindAllPaths(antFarm)
	fmt.Printf("len(paths): %v\n", len(paths))
}

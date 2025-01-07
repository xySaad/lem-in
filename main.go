package main

import (
	"fmt"
	"os"

	"lem-in/utils"
	"lem-in/utils/parser"
	"lem-in/utils/path_finder"
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
	fmt.Printf("path_finder.FindPaths(antFarm): %v\n", path_finder.FindPaths(antFarm))
	paths := utils.FindPaths(antFarm)
	fmt.Printf("paths: %v\n", paths)
	// fmt.Printf("len(paths): %v\n", len(paths))
	// for startLink, path := range paths {
	// 	fmt.Printf("startLink: %v ", startLink)
	// 	fmt.Printf("path: %v\n", path)
	// }
	// utils.DistributeAnts(antFarm, utils.ConvertPaths(antFarm, paths))
}

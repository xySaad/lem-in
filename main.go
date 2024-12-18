package main

import (
	"fmt"
	"os"

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

	fmt.Println("start:", antFarm.StartRoom, "end:", antFarm.EndRoom)
	for name, room := range antFarm.Rooms {
		fmt.Print("room:", name, " x:", room.X, " y:", room.Y, " links:", room.Links, "\n")
	}

	// paths := utils.FindAllPaths(antFarm)
	// fmt.Println(paths)
	// paths := bfs.Bfs(antFarm)
	// fmt.Println("paths from start to end")
	// for _, path := range paths {
	// 	fmt.Println(path)
	// }
}

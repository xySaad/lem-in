package main

import (
	"fmt"
	"lem-in/utils/parser"
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
	err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
}

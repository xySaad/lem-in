//go:build debug
// +build debug

package utils

import (
	"fmt"
)

func debugPrint(v ...interface{}) {
	fmt.Println(v...)
}

func debugPrintf(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}
func DebugPrint(v ...interface{}) {
	fmt.Println(v...)
}

func DebugPrintf(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}

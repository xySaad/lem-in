//go:build debug
// +build debug

package utils

import (
	"fmt"
)

func DebugPrint(v ...interface{}) {
	fmt.Println(v...)
}

func DebugPrintf(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}

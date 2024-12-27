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

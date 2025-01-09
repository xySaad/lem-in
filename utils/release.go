//go:build !debug
// +build !debug

package utils

func DebugPrint(v ...interface{}) {
	// No-op
}

func DebugPrintf(f string, v ...interface{}) {
	// No-op
}

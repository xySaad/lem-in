//go:build !debug
// +build !debug

package utils

func debugPrint(v ...interface{}) {
	// No-op
}

func debugPrintf(f string, v ...interface{}) {
	// No-op
}

package test

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// PrintStackTrace prints the stack trace of a panic.
func PrintStackTrace(testName string, recoverVal any) {
	var s strings.Builder
	s.WriteString("******************\n")
	s.WriteString(testName)
	s.WriteString(" panicked:\n")
	s.WriteString(fmt.Sprintf("%v", recoverVal))
	s.WriteString("\n\nStack trace from panic:\n")
	s.WriteString(string(debug.Stack()))
	s.WriteString("******************\n")
	fmt.Println(s.String())
}

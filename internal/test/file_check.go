package test

import (
	"os"
	"strings"
	"testing"
)

// FileToCheck returns a file name that exists and can be checked;
// if a matching _sol.md file exists, then that is chosen; otherwise the given file is used.
func FileToCheck(t *testing.T, questionFile string) string {
	t.Helper()
	questionFile = strings.ReplaceAll(questionFile, ".md", "")
	fileToCheck := questionFile + "_sol.md"
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		fileToCheck = questionFile + ".md"
	}
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		t.Error(err)
	}
	return fileToCheck
}

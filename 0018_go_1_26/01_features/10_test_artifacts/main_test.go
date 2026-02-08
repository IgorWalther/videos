package main

import (
	"os"
	"path/filepath"
	"testing"
)

// go test -v -artifacts -outputdir=./tmp
func TestFunc(t *testing.T) {
	dir := t.ArtifactDir()
	logFile := filepath.Join(dir, "tmp.log")
	content := []byte("ERROR")
	os.WriteFile(logFile, content, 0644)
	t.Log("Saved logs")
}

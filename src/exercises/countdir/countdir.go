package main

import (
	"io"
	"os"
)

// CountLines - count the lines in from Reader
func CountLines(r io.Reader) (int, error) {
	return 0, nil
}

// CountFile write to stdout number of lines in the file
func CountFile(path string) {
}

// CountDir - write to stdout numbers of lines for each files in directory
func CountDir(dir string) {
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}
}

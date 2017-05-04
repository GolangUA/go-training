package main

import (
	"io"
	"os"
)

// CountLines count lines in the reader
func CountLines(r io.Reader) (int, error) {
	return 0, nil
}

// CountFile write count of lines in the file to the stdout
func CountFile(path string) {
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountFile(arg)
	}
}

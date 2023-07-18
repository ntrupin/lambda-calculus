package misc

import (
	"fmt"
	"os"
)

// Open a file, panic on failure
func OpenFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
	return file
}

// Close a file, panic on failure
func CloseFile(file *os.File) {
	if file == nil {
		return
	}
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Length of cmd args != 2. Try again...\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s\n", filename)
		os.Exit(1)
	}

	byteCount, lineCount, wordCount := 0, 0, 0
	byteCount = len(data)

	var prev byte = 0
	for _, curr := range data {
		if curr == byte('\n') {
			lineCount++
		}

		if prev <= 0x20 && curr > 0x20 {
			wordCount++
		}

		prev = curr
	}

	fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, filename)
}

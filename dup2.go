package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 { // If there are no file names in the CL argument
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // os.Open() returns two values. File name and error.
			if err != nil {        // If there is an error opening the file.
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts) // If file name found
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) { // Count from file
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

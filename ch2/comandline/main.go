package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	// everything in command line argument
	// other than the first one.
	files := os.Args[1:]
	if len(files) < 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, count := range counts {
		if line != "" && count > 1 {
			fmt.Fprintf(os.Stdout, "%s repeated : %d times\n", line, count)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Err() == io.EOF{
			break
		}
		counts[scanner.Text()]++
	}
}

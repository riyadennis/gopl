package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// everything in command line argument
	// other than the first one.
	files := os.Args[1:]
	if len(files) < 0 {
		fmt.Fprint(os.Stderr, "no files to read")
		return
	} else {
		for _, file := range files {
			f, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			for _, l := range strings.Split(string(f), "\n"){
				counts[l]++
			}
		}
	}
	for line, count := range counts {
		if line != "" && count > 1 {
			fmt.Fprintf(os.Stdout, "%s repeated : %d times\n", line, count)
		}
	}
}


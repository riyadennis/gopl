package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type fileDetails struct {
	name string
	counts map [string]int
}

func main() {
	var fds []*fileDetails
	counts := make(map [string]int)
	// everything in command line argument
	// other than the first one.
	files := os.Args[1:]
	if len(files) < 0 {
		fmt.Fprint(os.Stderr, "no files to read")
		return
	} else {
		for _, file := range files {
			fd := &fileDetails{
				name:   file,
				counts: counts,
			}
			f, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			for _, l := range strings.Split(string(f), "\n"){
				fd.counts[l]++
			}
			fds = append(fds, fd)
		}
	}
	for _,  fd := range fds {
		for line, count := range fd.counts{
			if line != "" && count > 1 {
				fmt.Fprintf(os.Stdout, "%s repeated : %d times in file %s\n",
					line, count, fd.name)
			}
		}
	}
}


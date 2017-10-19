// dup4 prints count, text, and source of lines appearing more than once from list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	name  string
	files map[string]int
}

func main() {
	sources := make(map[Line]int)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		lineSources(f, arg, sources)
		f.Close()
	}

	for line, filenames := range sources {
		fmt.Printf("%s:\t", line)
		for filename, count := range filenames {
			fmt.Printf("%s - %d, ", filename, count)
		}
		fmt.Printf("\n")
	}
}

func lineSources(f *os.File, filename string, sources map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		sources[input.Text()][filename]++
	}
}

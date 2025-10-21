package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for filename, counter := range counts {
		for line, count := range counter {
			if count > 1 {
				fmt.Printf("%s\t%s\t%d\n", filename, line, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	filename := f.Name()

	if counts[filename] == nil {
		counts[filename] = make(map[string]int)
	}
	for input.Scan() {
		counts[filename][input.Text()]++
	}
}

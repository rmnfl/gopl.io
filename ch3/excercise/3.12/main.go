package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := os.Args[1]
	b := os.Args[2]

	fmt.Printf("Two strings are anagrams - %v\n", anagram(a, b))
}

func anagram(a, b string) bool {
	a = strings.ToLower(strings.ReplaceAll(a, " ", ""))
	b = strings.ToLower(strings.ReplaceAll(b, " ", ""))

	if len(a) != len(b) {
		return false
	}
	if a == b {
		return false
	}

	counts := make(map[rune]int)
	for _, ch := range a {
		counts[ch]++
	}
	for _, ch := range b {
		counts[ch]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}

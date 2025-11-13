package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	var frac string

	// getting optional sign
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	if dotPosition := strings.Index(s, "."); dotPosition != -1 {
		frac = s[dotPosition:]
		s = s[:dotPosition]
	}

	var i int
	if i = len(s) % 3; i == 0 {
		i = 3
	}

	for len(s) > 3 {
		buf.WriteString(s[:i])
		buf.WriteString(",")
		s = s[i:]
		i = 3
	}

	buf.WriteString(s)
	buf.WriteString(frac)

	return buf.String()
}

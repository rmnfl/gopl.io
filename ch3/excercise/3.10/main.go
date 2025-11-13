package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
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

	return buf.String()
}

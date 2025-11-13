package main

import (
	"fmt"
)

const (
	KB float64 = 1000
	MB         = KB * 1000
	GB         = MB * 1000
	TB         = GB * 1000
	PB         = TB * 1000
	EB         = PB * 1000
	ZB         = EB * 1000
	YB         = ZB * 1000
)

func main() {
	fmt.Printf("KB: %v\nMB: %v\nGB: %v\nTB: %v\nPB: %v\nEB: %v\nZB: %v\nYB: %v\n", KB, MB, GB, TB, PB, EB, ZB, YB)
}

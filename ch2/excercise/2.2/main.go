package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/excercise/2.2/lengthconv"
	"gopl.io/ch2/excercise/2.2/tempconv"
	"gopl.io/ch2/excercise/2.2/weightconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s number ...\n", os.Args[0])
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// Temperature conversions
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("Temperature:\n")
		fmt.Printf("\t%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("\t%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("\t%s = %s = %s\n", k, tempconv.KToC(k), tempconv.KToF(k))

		// Length conversions
		m := lengthconv.Meter(t)
		ft := lengthconv.Foot(t)
		fmt.Printf("Length:\n")
		fmt.Printf("\t%s = %s\n", m, lengthconv.MToFt(m))
		fmt.Printf("\t%s = %s\n", ft, lengthconv.FtToM(ft))

		// Weight conversions
		kg := weightconv.Kilogram(t)
		lb := weightconv.Pound(t)
		fmt.Printf("Weight:\n")
		fmt.Printf("\t%s = %s\n", kg, weightconv.KgToLb(kg))
		fmt.Printf("\t%s = %s\n", lb, weightconv.LbToKg(lb))

		fmt.Println("------------------------------------------------------------")
	}
}

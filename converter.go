package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use the command: converter <value> <type>")

		os.Exit(1)
	} else {
		fromUnit := os.Args[len(os.Args) - 1]
		values := os.Args[1 : len(os.Args) - 1]

		var toUnit string
		
		if fromUnit == "celsius" || fromUnit == "c" {
			toUnit = "fahrenheit"
		} else if fromUnit == "kilometers" || fromUnit == "k" {
			toUnit = "miles"
		} else {
			fmt.Printf("%s is an invalid unit.\n", fromUnit)
			os.Exit(1)
		}

		for i, v := range values {
			fromValue, err := strconv.ParseFloat(v, 64)

			if err != nil {
				fmt.Printf("The value %s in index %d is invalid.\n", v, i)

				os.Exit(1)
			} else {
				var toValue float64

				if toUnit == "fahrenheit" {
					fromUnit = "celsius"
					toValue = fromValue * 1.8 + 32
				} else {
					fromUnit = "kilometers"
					toValue = fromValue / 1.60934
				}

				fmt.Printf("%.2f %s => %.2f %s\n", fromValue, fromUnit, toValue, toUnit)
			}
		}

		os.Exit(0)
	}
}
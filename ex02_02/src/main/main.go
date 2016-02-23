package main

import (
	"fmt"
	"os"
	"strconv"
	"unitconv"
)

type UnitFunction func(value float64) string

func main() {
	var values []float64
	functions := []UnitFunction{
		unitconv.WeightUnits,
		unitconv.LengthUnits,
		unitconv.VelocityUnits,
		unitconv.ForceUnits,
		unitconv.PressureUnits,
	}

	if len(os.Args) == 1 {
		var tmp float64
		fmt.Print("value >> ")
		fmt.Scan(&tmp)
		values = append(values, tmp)
	} else {
		for _, arg := range os.Args {
			v, err := strconv.ParseFloat(arg, 64)

			if err == nil {
				values = append(values, v)
			}
		}
	}

	for _, v := range values {
		printUnits(functions, v)
	}
}

func printUnits(functions []UnitFunction, value float64) {
	fmt.Printf("<Value = %f>\n", value)
	for _, function := range functions {
		fmt.Println(function(value))
	}
}

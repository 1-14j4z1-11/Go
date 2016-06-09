package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelven float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Celsius = 0
	ctok          float64 = 273.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.5g C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.5g F", f)
}

func (k Kelven) String() string {
	return fmt.Sprintf("%.5g K", k)
}

func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) ToK() Kelven {
	return Kelven(float64(c) + ctok)
}

func (f Fahrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) ToK() Kelven {
	return f.ToC().ToK()
}

func (k Kelven) ToC() Celsius {
	return Celsius(float64(k) - ctok)
}

func (k Kelven) ToF() Fahrenheit {
	return k.ToC().ToF()
}

//////////////////////////////////////////////////////////////

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = Fahrenheit(value).ToC()
		return nil
	case "K":
		f.Celsius = Kelven(value).ToC()
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

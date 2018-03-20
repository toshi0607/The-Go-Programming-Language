package tempconv

import (
	"fmt"

	"github.com/prometheus/common/promlog/flag"
)

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Farenheit(value))
		return nil
	}
	return fmt.Errorf("invalid tempareture %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	boilingFreezingDiffF := (BoilingC - FreezingC).ToFahrenheit()
	fmt.Printf("%v\n", boilingFreezingDiffF)
}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (f Fahrenheit) ToCelsius() Celsius    { return Celsius((f - 32) * 5 / 9) }

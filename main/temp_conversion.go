package main

import "fmt"

type Fahrenheit float64
type Celsius float64

func (self Celsius) toString() string{
	return fmt.Sprintf("self %g ", self)
}

var f Fahrenheit
var c Celsius

func main() {
	f = 10
	c = 10
	fmt.Println("Fahrenheit of 30 degree is", ctf(30))
	fmt.Println("Fahrenheit of 30 degree is", c == Celsius(f))
	fmt.Println("Fahrenheit of 30 degree is", c.toString())
}

func ctf(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func ftc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

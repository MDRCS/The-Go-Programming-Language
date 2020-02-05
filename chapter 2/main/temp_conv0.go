package main

import (
	"fmt"
)

//type Celsius float64
//type Fahrenheit float64
type Kalvin float64

const (
	AbsolutZeroC Celsius = -273.15
	FreezingC Celsius = 0.0
	BoilingC Celsius = 100
)


func main() {

	fmt.Println(cTof(AbsolutZeroC))
	fmt.Println(cTof(BoilingC))

}

func cTof(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 +32)
}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5 / 9)
}

func cTok(c Celsius) Kalvin {
	return Kalvin(c)
}

func fTok(f Fahrenheit) Kalvin {
	return Kalvin((f - 32)*5 / 9)
}
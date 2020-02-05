package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

func celsToFah(c Celsius) Fahrenheit {
	return Fahrenheit((c*9 / 5) - 32)
}

func FahToCels(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5/9)
}

func main() {

	t,err := strconv.ParseFloat(os.Args[1],64)

	if err != nil {
		fmt.Fprintf(os.Stderr,"Parsing Error : %v",err)
		os.Exit(1)
	}

	conv := os.Args[2]

	if conv == "cels" {
		f := Fahrenheit(t)
		fmt.Printf("From Fahrenheit to Celsius : %v°F -> %v°C \n",f,FahToCels(f))
	}else
	if conv == "fah" {
		c := Celsius(t)
		fmt.Printf("From Celsius to Fahrenheit : %v°C -> %v°F \n",c,celsToFah(c))
	}


}
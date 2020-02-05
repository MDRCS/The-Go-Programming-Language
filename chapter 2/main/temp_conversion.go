package main


import (
	"fmt"
)

func main() {
var ffreezing_temp, fboiling_temp = 32.0,212.0

fmt.Printf("Freezing temperature : %g°F or %g°C \n",ffreezing_temp,fToc(ffreezing_temp))
fmt.Printf("Boiling temperature : %g°F or %g°C",fboiling_temp,fToc(fboiling_temp))

}

func fToc(f float64) float64 {
	return (f - 32)*5 / 9
}

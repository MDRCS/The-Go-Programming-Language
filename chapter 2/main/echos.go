package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","Separator")



func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))

	if !*n {
		fmt.Println()
	}

	primes := []int{2, 3, 5, 7, 11, 13}
	medals := []string{"gold", "silver", "bronze"}

	fmt.Println(medals,primes)
}
//
//$ go build gopl.io/ch2/echo4
//$ ./echo4 a bc def
//a bc def
//$ ./echo4 -s / a bc def
//a/bc/def
//$ ./echo4 -n a bc def
//a bc def$
//$ ./echo4 -help
//Usage of ./echo4:
//-n    omit trailing newline
//-s string
//separator (default " ")

package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		} }
}
func fib(x int) int {
	if x < 2 {
		return x }
	return fib(x-1) + fib(x-2)
}

// Notice how the program is expressed as the composition of two autonomous activities,
// spinning and Fibonacci computation. Each is written as a separate
// function but both make progress concurrently.
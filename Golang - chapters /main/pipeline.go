package main

import "fmt"

func main() {
	n := 10
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func(n int) {
		for x := 0; x<n ; x++ {
			naturals <- x
		}
		close(naturals)
	}(n)

	// Squarer
	go func() {
		for {
			x,ok := <-naturals
			if !ok {
				break
			}

			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// counts is a map -> (key : string ,value : int), a map is a data structure that provide store,fetch,delete,search in a constant time.
	// make() allows us to create new empty map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == " " {
			break
		}
		counts[input.Text()]++

	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			// fmt.Printf for formatted output
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

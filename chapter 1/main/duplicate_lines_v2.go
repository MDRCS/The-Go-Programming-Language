package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) < 1 {
		countlines(os.Stdin,counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplicate : %v",err)
				continue
			}

			countlines(f,counts)
			f.Close()
		}

		for line, occ := range counts {
			if occ > 1 {
				fmt.Printf("Number of occurence of each line | occ : %d , line : %s",occ,line)
			}
		}

	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if input.Text() == " " {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

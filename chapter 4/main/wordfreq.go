package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var wordfreq = make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == " " {
			break
		}

	words := strings.Fields(input.Text())
	for _,w := range words {
		wordfreq[w]++
		}

	err := input.Err()
	if err != nil {
			fmt.Fprintf(os.Stderr,"Error new scanner : %v",err)
		}

	}

	fmt.Println(wordfreq)
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)


// Normal loop

func main1 (){
	var s,sep string
	os.Args = []string{}
	os.Args = append(os.Args, "Mamadou")
	for i:= 1; i < len(os.Args); i++ {
		s+= sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

// Iteration

func main2 () {

	var s,sep string
	os.Args = []string{}
	os.Args = append(os.Args, "Mamadou")
	for _,arg := range os.Args[1:] {
		//fmt.Println(key)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

// Using join to make all the arguments as a string and it has the best execution time

func main3 () {
	os.Args = []string{}
	os.Args = append(os.Args, "Mamadou")
	fmt.Println(strings.Join(os.Args[1:]," "))
}

func main () {
	// Testing the most effcient implementation to concatenate args and display them
	start := time.Now()
	main1()
	fmt.Println("Execution time for main1() -> ",time.Now().Sub(start))

	start = time.Now()
	main2()
	fmt.Println("Execution time for main2() -> ",time.Now().Sub(start))

	start = time.Now()
	main3()
	fmt.Println("Execution time for main3() -> ",time.Now().Sub(start))

	// Result :
	// Mamadou
	//Execution time for main1() ->  28.575µs
	//Mamadou
	//Execution time for main2() ->  2.453µs
	//Mamadou
	//Execution time for main3() ->  2.278µs
}

// Command -> go run command_line_args.go HELLO MAMADOU
// Result -> HELLO MAMADOU


package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	person := make(map[string]int)

	person["med"] = 24
	person["aymen"] = 25
	person["anas"] = 23

	//fmt.Println(person["med"])

	delete(person,"anas")

	for name, age := range person {
		fmt.Printf("%s \t%d\n", name, age)
	}

	//var names []string
	names := make([]string,0,len(person))

	for name :=range person {
		names = append(names,name)
	}

	sort.Strings(names)

	for _,name := range names {
		fmt.Printf("%s \t%d\n",name,person[name])
	}

	// checking if the element exist or not

	age,ok := person["abdo"]

	if !ok {
		fmt.Fprintf(os.Stderr,"Element doesn't exist \n")
		os.Exit(1)
	}

	fmt.Println(age)


}
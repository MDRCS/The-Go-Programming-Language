package main

import "fmt"

func sum(vals ...int) int {
	var s = 0
	for _,val := range vals {
		s+=val
	}

	return s
}

func min(values ...int) {
	min := values[0]

	for _,value := range values {
		if value < min {
			min,value = value,min
		}
	}
}

func max(values ...int) int{
	max := values[0]

	for _,value := range values {
		if value > max {
			max,value = value,max
		}
	}

	return max
}

func main() {
	var vals = []int{1, 2, 3, 4, 5}
	var s = sum(vals...)
	fmt.Println(s)
	fmt.Println(max(vals...))

}

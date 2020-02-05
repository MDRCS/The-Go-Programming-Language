package main

import (
	"fmt"
)

func main() {

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := []string{USD: "$", EUR: "9", GBP: "!", RMB: "*"}

	fmt.Printf("%d%s\n",RMB,symbol[RMB])

	medals := []string{"Gold","Bronze","Silver"}

	first := medals[:1] //Slice

	fmt.Println(first,cap(first))

	weekday := [...]int{1,2,3,4,5,6,7}
	fmt.Println(weekday)
	reverse_ptr(&weekday)
	fmt.Println(weekday)
	//rotate_left(2,weekday)
	//
	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%d \t%v\n", i, cap(y), y)
	//	x=y
	//}
}

func reverse_ptr(s *[7]int) {
	j:= len(s)-1
	i:= 0
	for i<j {
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
}

func reverse(s []int) {
	j:= len(s)-1
	i:= 0
	for i<j {
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
}

// Rotate s left by two positions.
func rotate_left(n int,list []int) {

	if n<len(list){

		reverse(list[:n])
		reverse(list[n:])
		reverse(list)
		fmt.Println(list)
	}

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int,zlen,zcap)
		copy(z,x)
	}

	z[len(x)] = y
	return z

}

func nonempty(list []string){

	i := 0
	for _,s := range list {

		if s != "" {
			list[i] = s
			i++
		}
	}

}
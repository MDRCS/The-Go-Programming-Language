package main

import (
	"fmt"
)

func main() {
	var x = 1
	var p = &x //address of x variable
	fmt.Printf("value of the p variable is %d == %d \n",*p,x) //1 1
	*p = 2 //assign p with a new value
	fmt.Printf("the new value of x is %d \n",x) //2
	var d = f()
	fmt.Printf("%v \n",d)
	var count = 0
	fmt.Printf("Before update value of the counter : %d \n ",count)
	inc(&count)
	fmt.Printf("After update value of the counter : %d \n",count)

}

func inc(va *int) *int { //I specified that i am going to return address of a variable *int and not a value
	*va++
	return va
}

func f() *int {
	v:=1
	return &v
}
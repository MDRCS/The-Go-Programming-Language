package main

import (
	"bytes"
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

func main() {
	// "program" in Japanese katakana
	s := ">+=@?"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r)) // ">+=@?"
	fmt.Println(string(0x4eac))
	fmt.Println(string(s))
	fmt.Println(string(1234567))
	fmt.Println(comma("12345679"))

	////////////////////////////////////////////

	fmt.Println(GiB)
}


func comma(value string) string {

	var buf bytes.Buffer

	for i,digit :=range value {
		buf.WriteString(string(digit))
		if i%3 == 0 {
			buf.WriteString(",")
		}
	}

	return buf.String()

}
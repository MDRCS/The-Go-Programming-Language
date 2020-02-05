package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	//crypto_hash()

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func crypto_hash() {
	string_hash := os.Args[1]

	if len(os.Args) > 2 {
		hash_func := os.Args[2]

		if hash_func == "sha512" {
			hash := sha512.Sum512([]byte(string_hash))
			fmt.Printf("Hash of %s is %x using sha512\n",string_hash,hash)
		} else
		if hash_func == "sha384" {
			hash := sha512.Sum384([]byte(string_hash))
			fmt.Printf("Hash of %s is %x using sha384\n",string_hash,hash)
		}

		return
	}

	hash := sha256.Sum256([]byte(string_hash))
	fmt.Printf("Hash of %s is %x using sha256\n",string_hash,hash)

	//MACBOOK-ELRAHALI:main mdrahali$ go run sha256.go "Mohamed EL Rahali"
	//Hash of Mohamed EL Rahali is 96c6905917c811b8e384dd2f6ac4e5263410d5efd1f0af7718e6bd9bce9bb98f using sha256
}
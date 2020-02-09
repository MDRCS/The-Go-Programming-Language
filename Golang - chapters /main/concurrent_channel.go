package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func mirroredQuery() []byte {
	responses := make(chan []byte, 3)

	go func() {
		resp,err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			log.Fatal(err)
		}
		response,err := json.Marshal(resp.Body)
		if err!=nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.Body)
		fmt.Println(response)
		fmt.Println(string(response))
		responses <- response
	}()

	go func() {
		resp,err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			log.Fatal(err)
		}
		response,err := json.Marshal(resp.Body)
		if err!=nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.Body)
		fmt.Println(response)
		fmt.Println(string(response))
		responses <- response
	}()

	go func() {
		resp,err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			log.Fatal(err)
		}
		response,err := json.Marshal(resp.Body)
		if err!=nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.Body)
		fmt.Println(response)
		fmt.Println(string(response))
		responses <- response
	}()

	return <-responses // return the quickest response
}

func main() {
	// get the quuickest response
	s := mirroredQuery()

	fmt.Println(s)
}

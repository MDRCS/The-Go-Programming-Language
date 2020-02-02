package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	response,err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fetch Error : %v",err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr,"Reda body error : %v",err)
		os.Exit(1)
	}

	fmt.Printf("%s",body)

}

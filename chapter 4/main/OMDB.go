package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const OMDB_API = "http://www.omdbapi.com/?t=%s&apikey=a38c58d0"
const ImagePath = "../img/%s.jpg" //for command-line use this path else delete one dot

type Movie_ struct {
	Title string
	Year string
	Awards string
	Poster string
	Ratings []*Rating
}

type Rating struct {
	Source string
	Value string
}

func search_movie(movie string) (*Movie_,error) {
	url := fmt.Sprintf(OMDB_API,movie)

	response,err := http.Get(url)
	if err!= nil {
		log.Fatalf("Error : %s",err)
		return nil,err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", response.Status)
	}

	var result Movie_
	if err := json.NewDecoder(response.Body).Decode(&result) ; err != nil {
		response.Body.Close()
		return nil,err
	}

	response.Body.Close()
	return &result,nil

}


func poster(imageUrl string,movie string) {

	// don't worry about errors
	path := fmt.Sprintf(ImagePath,movie)
	response, e := http.Get(imageUrl)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	fmt.Println(path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}

func main() {
	var movie = "Interstellar"

	//var movie = os.Args[1]
	result,err := search_movie(movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Movie description : \n %s, %s, %s \n \n",result.Title, result.Year, result.Awards)
	fmt.Println("Ratings : \n")
	for _,ratings := range result.Ratings {
		fmt.Printf("%s, %s \n",ratings.Source,ratings.Value)
	}

	//Download image
	poster(result.Poster,movie)
}
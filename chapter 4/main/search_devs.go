package main

// check github api documentation for more : https://help.github.com/en/github/searching-for-information-on-github/searching-users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const github_api = "https://api.github.com/search/users"

type Github_search_result struct {
	TotalCount int `json:"total_count"`
	Items []*Info
}

type Info struct {
	Login string
	Url string `json:"html_url"`
	Subscription string `json:"subscriptions_url"`
	Score float64
}

func main() {

	result,err := search_user(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of devs : %d \n",result.TotalCount)
	for _, item :=range result.Items {
		//fmt.Println(item.Url)
		fmt.Printf(" %s %.55s %9.2f\n", item.Url,item.Login,item.Score)
	}

}


func search_user(terms []string) (*Github_search_result,error){
	q := url.QueryEscape(strings.Join(terms," "))
	response,err := http.Get(github_api + "?q=" + q + "&order=desc")
	if err != nil {
		return nil,err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil,fmt.Errorf("Search query failed : %s",err)
	}

	var result Github_search_result
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()
		return nil,err
	}

	response.Body.Close()
	return &result,nil

}
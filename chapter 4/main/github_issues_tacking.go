package main
// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
// Also see https://developer.github.com/v3/issues/#create-an-issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const PostIssue = "https://github.com/%s/%s/issues" //we can create issue for any repo


type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type CreateIssue struct {
	Title string
	Body string
	assignees []string
	milestone int
	labels []string
}

//Post an Issue
func CreateIssue_(issue CreateIssue, owner string,repo string) string {

	url := fmt.Sprintf(PostIssue, owner, repo)
	marshalBody, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("create issue: error marshaling json - %v\n", err)
	}

	response, err := http.Post(url, "application/vnd.github.symmetra-preview+json", bytes.NewBuffer(marshalBody))

	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatalf("create issue: response from github API %d\n", response.StatusCode)
	}

	fmt.Println(response.Body)

	var result IssuesSearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		response.Body.Close()
	}
	response.Body.Close()

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.HTMLURL)
	}

	return url
}

// SearchIssues queries the GitHub issue tracker."
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//curl https://api.github.com/search/issues?q=windows+label:bug+language:python+state:open&sort=created

func main() {
	// command-line go run github.go is:open json decoder
	// go run github.go windows label:bug language:python state:open
	//result, err := SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s %s\n",
	//		item.Number, item.User.Login, item.Title,item.HTMLURL)
	//}


	//Post Issue

	var owner = "MDRCS"
	var repo = "The-Go-Programming-Language"
	var assignees = []string{owner}
	var labels = []string{"labels"}
	var issue = CreateIssue{"Found a bug","I'm having a problem with this.",assignees,1,labels}
	CreateIssue_(issue,owner,repo)

}

//{
//"title": "Found a bug",
//"body": "I'm having a problem with this.",
//"assignees": [
//"octocat"
//],
//"milestone": 1,
//"labels": [
//"bug"
//]
//}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Issue struct {
	Url   string `json:"html_url"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SearchResult struct {
	TotalCount int     `json:"total_count"`
	Items      []Issue `json:"items"`
}

type Link struct {
	URL  string
	Type string
}

func ParseLink(header string) []Link {
	links := []Link{}
	for _, v := range strings.Split(header, ",") {
		l := Link{}
		trimed := strings.Trim(v, " ")
		for _, t := range strings.Split(trimed, ";") {
			if t[0] == '<' && t[len(t)-1] == '>' {
				l.URL = strings.Trim(t, "<>")
				continue
			}
			if t[1] == 'r' && t[4] == '=' {
				s := strings.Split(t, "=")
				l.Type = strings.Trim(s[1], "\"")
				continue
			}
		}
		links = append(links, l)
	}
	return links
}

func GetIssues(url, token string) []Issue {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Issue{}
	}
	req.Header.Add("Authorization", "token "+token)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []Issue{}
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	var result SearchResult
	dec.Decode(&result)

	links := ParseLink(res.Header.Get("Link"))
	for _, v := range links {
		if v.Type == "next" {
			i := GetIssues(v.URL, token)
			return append(i, result.Items...)
		}
	}

	return result.Items
}

func main() {
	username := ""
	token := ""
	args := os.Args
	if len(args) < 2 {
		fmt.Println("failed")
		os.Exit(1)
	}
	username = args[1]
	token = args[2]
	v := url.Values{}
	v.Add("q", fmt.Sprintf("involves:%s closed:2018-01-01..2018-12-31", username))
	v.Add("per_page", "100")
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/search/issues", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.URL.RawQuery = v.Encode()
	issues := GetIssues(req.URL.String(), token)

	for _, v := range issues {
		fmt.Printf("[%s](%s)\n", v.Title, v.Url)
	}
}

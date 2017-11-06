package handler

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/gommon/log"
)

const GitHubSearchURL = "https://api.github.com/search/repositories"

type GithubApi interface {
	SearchRepositories(w http.ResponseWriter, r *http.Request)
}

func SearchRepositories(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.RawQuery
	query, err := url.ParseQuery(queries)
	if err != nil {
		log.Printf("parse query error. err: %v", err)
	}

	values := url.Values{}
	for k, v := range query {
		switch k {
		case "q":
			values.Add("q", v[0])
		}
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", GitHubSearchURL, nil)
	if err != nil {
		log.Fatalf("build request error. err: %v")
	}
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request error. err: %v", err)
	}
	if resp.Status != "200 OK" {
		log.Fatalf("http request failed. code: %v", resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatalf("read error. err: %v")
	}

	w.Write([]byte("Search Result:\n"))
	w.Write(b)
}

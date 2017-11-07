package service

import (
	"log"
	"net/http"
	"net/url"
)

const githubSearchURL = "https://api.github.com/search/repositories"

type GithubService interface {
	SearchRepositories(r *http.Request)
}

type GithubClient struct {
	GithubService
}

func (g GithubClient) New() http.Client {
	return http.Client{}
}

func (g GithubClient) SearchRepositories(r *http.Request) (*http.Response, error) {
	queries := r.URL.RawQuery
	query, err := url.ParseQuery(queries)
	if err != nil {
		log.Printf("parse query error. err: %v", err)
		return nil, err
	}

	values := url.Values{}
	for k, v := range query {
		switch k {
		case "q":
			values.Add("q", v[0])
		}
	}

	client := g.New()
	req, err := http.NewRequest("GET", githubSearchURL, nil)
	if err != nil {
		log.Fatalf("build request error. err: %v")
	}
	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request error. err: %v", err)
		return nil, err
	}
	if resp.Status != "200 OK" {
		log.Fatalf("http request failed. code: %v", resp.Status)
		return nil, err
	}

	return resp, nil
}

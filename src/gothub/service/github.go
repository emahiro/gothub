package service

import (
	"log"
	"net/http"
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
	client := g.New()
	req, err := http.NewRequest("GET", githubSearchURL, nil)
	if err != nil {
		log.Fatalf("build request error. err: %v", err)
	}
	// form をparseする
	if err := r.ParseForm(); err != nil {
		log.Fatalf("parse form error. err: %v", err)
		return nil, err
	}

	req.URL.RawQuery = r.PostForm.Encode()
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

package handler

import (
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
	values, err := url.ParseQuery(queries)
	if err != nil {
		log.Printf("parse query error. err: %v", err)
	}

	for k, v := range values {
		log.Printf("key: %v, val: %v", k, v)

	}
}

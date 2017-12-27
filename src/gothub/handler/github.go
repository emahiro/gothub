package handler

import (
	"encoding/json"
	"net/http"

	"gothub/model"
	"gothub/render"
	"gothub/service"

	"github.com/labstack/gommon/log"
)

func SearchRepositories(w http.ResponseWriter, r *http.Request) {
	render.RenderHTML("template/github/search_repositories.tmpl", w, nil)
}

func PostSearchRepositories(w http.ResponseWriter, r *http.Request) {
	c := service.GithubClient{}
	resp, err := c.SearchRepositories(r)
	if err != nil {
		log.Fatalf("request error. err: %v")
	}
	resp.Body.Close()

	var repos model.Repositories
	if json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatalf("json unmarshal error. err: %v", err)
	}

	json.NewEncoder(w).Encode(&repos)
}

package handler

import (
	"encoding/json"
	"io/ioutil"
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
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var repos model.Repositories
	if err := json.Unmarshal(b, &repos); err != nil {
		log.Fatalf("json unmarshal error. err: %v", err)
	}

	json.NewEncoder(w).Encode(&repos)
}

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
	github := service.GithubClient{}
	resp, err := github.SearchRepositories(r)
	if err != nil {
		log.Fatalf("request error. err: %v")
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var repos model.Repositories
	if err := json.Unmarshal(b, &repos); err != nil {
		log.Fatalf("json unmarshal error. err: %v", err)
	}

	render.RenderHTML("template/github/search.tmpl", w, repos)
}

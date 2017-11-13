package handler

import (
	"io/ioutil"
	"net/http"

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

	w.Write([]byte("Search Result:\n"))
	w.Write(b)

}

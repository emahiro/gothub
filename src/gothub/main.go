package main

import (
	"gothub/handler"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const port = "8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Top).Methods("GET")
	{
		// github request
		router.HandleFunc("/github/search/repositories", handler.SearchRepositories).Methods("GET")
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("err: %v", err)
	}
}

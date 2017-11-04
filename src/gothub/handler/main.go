package handler

import (
	"net/http"

	"fmt"

	"github.com/labstack/gommon/log"
)

func Top(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("%v", r.Body)
}

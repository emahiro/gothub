package handler

import (
	"net/http"

	"gothub/render"
)

func Top(w http.ResponseWriter, r *http.Request) {
	sampleData := struct {
		Name string
	}{
		Name: "taro",
	}
	render.RenderHTML("template/index.tmpl", w, sampleData)
}

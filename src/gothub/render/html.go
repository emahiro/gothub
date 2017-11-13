package render

import (
	"html/template"
	"net/http"

	"github.com/labstack/gommon/log"
)

func RenderHTML(f string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(f)
	if err != nil {
		log.Fatalf("parse template error. err: %v", err)
	}
	return t.Execute(w, data)
}

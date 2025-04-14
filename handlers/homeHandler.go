package handler

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed index.html
var indexHTML string

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	tmpl, err := template.New("index.html").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, struct {
		Path string
	}{
		Path: path,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

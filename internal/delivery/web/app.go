package web

import (
	"html/template"
	"net/http"
)

type app struct {
	tpl template.Template
}

func (app app) indexHandler(wr http.ResponseWriter, req *http.Request) {
	app.tpl.ExecuteTemplate(wr, "index.html", nil)
}

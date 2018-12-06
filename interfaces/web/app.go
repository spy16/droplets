package web

import (
	"html/template"
	"net/http"
)

type app struct {
	render func(wr http.ResponseWriter, tpl string, data interface{})
	tpl    template.Template
}

func (app app) indexHandler(wr http.ResponseWriter, req *http.Request) {
	app.render(wr, "index.tpl", nil)
}

package web

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/spy16/droplets/pkg/logger"
)

// New initializes a new webapp server.
func New(lg logger.Logger, cfg Config) (http.Handler, error) {
	tpl, err := initTemplate(lg, "", cfg.TemplateDir)
	if err != nil {
		return nil, err
	}

	app := &app{
		render: func(wr http.ResponseWriter, tplName string, data interface{}) {
			if err := tpl.ExecuteTemplate(wr, tplName, data); err != nil {
				lg.Errorf("failed to render template '%s': %+v", tplName, err)
			}
		},
	}

	fsServer := newSafeFileSystemServer(lg, cfg.StaticDir)

	router := mux.NewRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", fsServer))
	router.Handle("/favicon.ico", fsServer)

	// web app routes
	router.HandleFunc("/", app.indexHandler)

	return router, nil
}

// Config represents server configuration.
type Config struct {
	TemplateDir string
	StaticDir   string
}

func initTemplate(lg logger.Logger, name, path string) (*template.Template, error) {
	apath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(apath)
	if err != nil {
		return nil, err
	}

	lg.Infof("loading templates from '%s'...", path)
	tpl := template.New(name)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fp := filepath.Join(apath, f.Name())
		lg.Debugf("parsing template file '%s'", f.Name())
		tpl.New(f.Name()).ParseFiles(fp)
	}

	return tpl, nil
}

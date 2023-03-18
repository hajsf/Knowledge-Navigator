package gpt

import (
	"fmt"
	"hajsf/config"
	"hajsf/router"
	"html/template"
	"io/fs"
	"net/http"
)

var messages []map[string]string

var prompt string

func init() {
	router.R.Route(http.MethodGet, "/gpt", Page1Handler)
	router.R.Route(http.MethodPost, "/gpt/api/data", Postfunc)
	messages = make([]map[string]string, 0)
}

func Page1Handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "An error occurred: %v", r)
		}
	}()

	var t *template.Template

	if config.Debug {
		t = template.Must(template.ParseFiles("root/layout.html", "gpt/default.html"))
	} else {
		fsys := fs.FS(embeddedFiles)
		t = template.Must(template.ParseFS(fsys, "root/layout.html", "gpt/default.html"))
	}
	t.Execute(w, nil)
}

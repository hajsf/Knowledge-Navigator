package world

import (
	"fmt"
	"hajsf/router"
	"net/http"
	"text/template"
)

func init() {
	router.R.Route(http.MethodGet, "/world", WorldHandler)
}

func WorldHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "World!")
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "An error occurred: %v", r)
		}
	}()

	// Parse layout and view files
	t := template.Must(template.ParseFiles("root/layout.html", "world/world.html"))

	// Execute layout with view as data
	t.Execute(w, nil)
}

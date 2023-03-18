package hello

import (
	"fmt"
	"net/http"

	"hajsf/router"
)

func init() {
	router.R.Route(http.MethodGet, "/hello/:name", HelloHandler)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := router.URLParam(r, "name")
	fmt.Fprintf(w, "Hello %s!", name)
}

package root

import (
	"fmt"
	"html/template"
	"net/http"

	"hajsf/router"
)

func init() {
	// This function registers a new route with the router for handling HTTP GET requests to the "/page1" path.
	//
	// When a request is received at this path, the function first sets up a deferred function to recover from any panics that may occur during execution. If a panic occurs, an error message will be written to the response using fmt.Fprintf.
	//
	// Next, the function parses two template files: "layout.html" and "view1.html". These files define the layout and content of the page that will be returned in response to the request. The template.Must function is used to ensure that any errors that occur while parsing the templates cause a panic.
	//
	// Finally, the parsed templates are executed with nil data and written to the response. This causes the content of the templates to be rendered and sent back to the client as an HTML page.
	router.R.Route(http.MethodGet, "/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "An error occurred: %v", r)
		}
	}()

	// Parse layout and view files
	t := template.Must(template.ParseFiles("root/layout.html", "root/index.html"))

	// Execute layout with view as data
	t.Execute(w, nil)
}

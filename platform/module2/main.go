package module1

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"hajsf/router"
)

type contextKey string

const userIsAllowedKey contextKey = "userIsAllowed"

func init() {
	// This function registers a new route with the router for handling HTTP GET requests to the "/page1" path.
	//
	// When a request is received at this path, the function first sets up a deferred function to recover from any panics that may occur during execution. If a panic occurs, an error message will be written to the response using fmt.Fprintf.
	//
	// Next, the function parses two template files: "layout.html" and "view1.html". These files define the layout and content of the page that will be returned in response to the request. The template.Must function is used to ensure that any errors that occur while parsing the templates cause a panic.
	//
	// Finally, the parsed templates are executed with nil data and written to the response. This causes the content of the templates to be rendered and sent back to the client as an HTML page.
	router.R.Route(http.MethodGet, "/module2", checkUser(Page2Handler))
}

func checkUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// check if user is allowed
		userIsAllowed := false //or true // set this value based on your logic

		// add information to context
		ctx := context.WithValue(req.Context(), userIsAllowedKey, userIsAllowed)
		log.Println("Before calling next handler")
		next(w, req.WithContext(ctx))
		log.Println("After calling next handler")
	}
}

func Page2Handler(w http.ResponseWriter, r *http.Request) {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "An error occurred: %v", r)
		}
	}()

	// Parse layout and view files
	t := template.Must(template.ParseFiles("root/layout.html", "module2/view2.html"))

	// get information from context
	var userIsAllowed bool
	value := r.Context().Value(userIsAllowedKey)
	if value == nil {
		// Handle the case where the value is nil
		userIsAllowed = false
	} else {
		userIsAllowed = value.(bool)
		// Continue with your code
	}

	data := struct {
		UserIsAllowed bool
	}{
		UserIsAllowed: userIsAllowed,
	}
	t.Execute(w, data)
}

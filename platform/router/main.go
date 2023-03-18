// Package router provides a simple HTTP router implementation in Go.
//
// It allows you to define routes with a specific HTTP method and path pattern
// and associate them with a handler function. When a request comes in, the
// router will try to find a matching route based on the request's method and
// path. If it finds one, it will call the associated handler function and pass
// in any URL parameters that were extracted from the path.
//
// Example usage:
//
//	r := router.New()
//	r.Route("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    id := router.URLParam(r, "id")
//	    fmt.Fprintf(w, "User ID: %s\n", id)
//	})
//
//	http.ListenAndServe(":8080", r)
package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// R is a package-level instance of the Router type.
//
// It is initialized by the init function when the package is imported and can
// be used to define and handle routes without creating a new Router instance.
//
// Example usage:
//
//	router.R.Route("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    id := router.URLParam(r, "id")
//	    fmt.Fprintf(w, "User ID: %s\n", id)
//	})
var R *Router

// init initializes the package-level Router instance.
//
// This function is called automatically when the package is imported. It
// creates a new Router instance and assigns it to the package-level R variable.
func init() {
	// Get our router configuration
	R = new(Router)
}

// ctxKey is a string type used as a key for storing values in a context.Context.
//
// It is used to define the paramsKey constant, which is used as the key for
// storing URL parameters in the request context. This allows handler functions
// to access URL parameters using the URLParam function.
type ctxKey string

// paramsKey is a context key used for storing URL parameters in the request context.
//
// It is defined as a ctxKey type with the value "params". This allows handler
// functions to access URL parameters using the URLParam function by retrieving
// them from the request context using this key.
const paramsKey ctxKey = "params"

// RouteEntry represents a single route in the router.
//
// It contains the HTTP method, path pattern, and handler function for the route.
// The router uses this information to match incoming requests to the appropriate
// handler function based on their method and path.
type RouteEntry struct {
	Method      string
	Path        *regexp.Regexp
	HandlerFunc http.HandlerFunc
}

// Router is a simple HTTP router implementation in Go.
//
// It allows you to define routes with a specific HTTP method and path pattern
// and associate them with a handler function. When a request comes in, the
// router will try to find a matching route based on the request's method and
// path. If it finds one, it will call the associated handler function and pass
// in any URL parameters that were extracted from the path.
type Router struct {
	routes []*RouteEntry
}

// Route defines a new route with the specified method, path pattern, and handler function.
//
// The method parameter specifies the HTTP method for the route (e.g. "GET", "POST", etc.).
// The path parameter specifies the path pattern for the route using a simple syntax
// where named parameters are indicated by a leading colon (e.g. "/users/:id").
// The handler parameter specifies the function that will be called to handle requests
// that match this route.
//
// Example usage:
//
//	r := router.New()
//	r.Route("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    id := router.URLParam(r, "id")
//	    fmt.Fprintf(w, "User ID: %s\n", id)
//	})
func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	log.Printf("Adding route: %s %s", method, path)
	compiledPath := compilePath(path)
	routeEntry := &RouteEntry{
		Method:      method,
		Path:        compiledPath,
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, routeEntry)
}

// ServeHTTP implements the http.Handler interface for the Router type.
//
// It is called by the net/http package to handle incoming HTTP requests. The
// method will try to find a matching route based on the request's method and
// path. If it finds one, it will call the associated handler function and pass
// in any URL parameters that were extracted from the path.
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	for _, route := range rtr.routes {
		if route.Method != r.Method {
			continue
		}

		match := route.Path.FindStringSubmatch(r.URL.Path)
		if match == nil {
			continue
		}

		params := make(map[string]string)
		for i, name := range route.Path.SubexpNames() {
			if i == 0 || name == "" {
				continue
			}
			params[name] = match[i]
		}

		log.Printf("Matched route: %s %s", route.Method, route.Path.String())

		ctx := context.WithValue(r.Context(), paramsKey, params)
		route.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))

		return
	}

	log.Println("No matching route found")
	http.NotFound(w, r)
}

// URLParam retrieves the value of the specified URL parameter from the request context.
//
// The r parameter specifies the *http.Request from which to retrieve the parameter.
// The name parameter specifies the name of the parameter to retrieve.
//
// If a value for the specified parameter is found in the request context, it will be
// returned as a string. Otherwise, an empty string will be returned.
//
// Example usage:
//
//	r := router.New()
//	r.Route("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
//	    id := router.URLParam(r, "id")
//	    fmt.Fprintf(w, "User ID: %s\n", id)
//	})
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params, ok := ctx.Value(paramsKey).(map[string]string)
	if !ok {
		return ""
	}
	value, ok := params[name]
	if !ok {
		return ""
	}
	return value
}

// compilePath compiles a path pattern into a regular expression and extracts any named parameters.
//
// The path parameter specifies the path pattern to compile. It uses a simple syntax
// where named parameters are indicated by a leading colon (e.g. "/users/:id").
//
// The function returns two values: a *regexp.Regexp representing the compiled path
// pattern and a []string containing the names of any extracted parameters.
//
// Example usage:
//
//	re, params := compilePath("/users/:id")
//	fmt.Println(re.MatchString("/users/123")) // true
//	fmt.Println(params)                        // ["id"]
func compilePath(path string) *regexp.Regexp {
	var regex strings.Builder
	regex.WriteString(`^`)
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			regex.WriteString(`(?P<`)
			regex.WriteString(part[1:])
			regex.WriteString(`>[^/]+)`)
		} else if strings.HasPrefix(part, "*") {
			regex.WriteString(`(?P<`)
			regex.WriteString(part[1:])
			regex.WriteString(`>.+)`)
		} else {
			regex.WriteString(regexp.QuoteMeta(part))
		}
		regex.WriteString("/")
	}
	regexString := regex.String()[:regex.Len()-1]
	regexString += `$`
	return regexp.MustCompile(regexString)
}

/*
func compilePath(path string) *regexp.Regexp {
	var regex strings.Builder
	regex.WriteString(`^`)

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			regex.WriteString(`(?P<`)
			regex.WriteString(part[1:])
			regex.WriteString(`>[^/]+)`)
		} else {
			regex.WriteString(regexp.QuoteMeta(part))
		}
		regex.WriteString("/")
	}

	// Remove trailing slash from regex
	regexString := regex.String()[:regex.Len()-1]

	regexString += `$`
	return regexp.MustCompile(regexString)
}
*/

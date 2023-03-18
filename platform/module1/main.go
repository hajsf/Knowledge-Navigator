// Package module1 provides a simple web server that serves content using Go's html/template package.
package module1

// The import keyword is used to include packages in a Go source file. The packages are specified by their import paths and are separated by spaces within the parentheses. This line of code imports the "fmt" and "net/http" packages so that their exported identifiers can be used in this source file.
import (
	"embed" // This package provides access to files embedded in the program during compilation using the //go:embed directive.
	"fmt"   // This package implements formatted I/O with functions analogous to C’s printf and scanf.
	"hajsf/config"
	"hajsf/router"  // This is a custom package that is being imported. It’s our used for routing HTTP requests.
	"html/template" // This package implements data-driven templates for generating HTML output safe against code injection.
	"io/fs"         // This package provides a file system interface that can be implemented by other packages to provide access to different types of file systems.
	"net/http"      // This package provides HTTP client and server implementations.
)

// embeddedFiles is an embedded file system that contains the HTML templates used by the Page1Handler function.
//
//go:embed view1.html
var embeddedFiles embed.FS

// init registers a route with the router.R object to handle GET requests to the /module1 endpoint using the Page1Handler function.
func init() {
	// This function registers a new route with the router for handling HTTP GET requests to the "/page1" path.
	//
	// When a request is received at this path, the function first sets up a deferred function to recover from any panics that may occur during execution. If a panic occurs, an error message will be written to the response using fmt.Fprintf.
	//
	// Next, the function parses two template files: "layout.html" and "view1.html". These files define the layout and content of the page that will be returned in response to the request. The template.Must function is used to ensure that any errors that occur while parsing the templates cause a panic.
	//
	// Finally, the parsed templates are executed with nil data and written to the response. This causes the content of the templates to be rendered and sent back to the client as an HTML page.
	router.R.Route(http.MethodGet, "/module1", Page1Handler)
}

// Page1Handler is an http.HandlerFunc that serves content using Go's html/template package. If the debug variable is true, this function reads template files from disk. Otherwise, it reads template files from the embedded file system.
func Page1Handler(w http.ResponseWriter, r *http.Request) {
	// This defer statement sets up a function to be called when the surrounding function returns. The deferred function is used to recover from any panics that may occur during execution of the surrounding function.
	//
	// If a panic occurs, the recover function will be called and will return the value that was passed to the panic function. In this case, we are using the value returned by recover to print an error message to the response using fmt.Fprintf.
	//
	// This allows us to handle panics gracefully and keep our app running even if an error occurs. The error message provides some information about what went wrong so that we can diagnose and fix the issue.
	defer func() {
		// The recover function is used to regain control of a panicking goroutine. If the deferred function calls recover and a panic has occurred in the same goroutine, recover returns the value passed to panic. Otherwise, recover returns nil. This line of code checks if a panic has occurred by calling the recover function and assigning its return value to the r variable. If r is not nil, this means that a panic has occurred and appropriate action can be taken.
		if r := recover(); r != nil {
			// The fmt.Fprintf function is used to write formatted output to an io.Writer. The first argument is the io.Writer to write to, the second argument is a format string that specifies how the remaining arguments should be formatted, and the remaining arguments are the values to format. This line of code writes a formatted error message to the w io.Writer that includes the value of the r variable.
			fmt.Fprintf(w, "An error occurred: %v", r)
		}
		// This line of code defines and immediately invokes an anonymous function. The function body is enclosed in curly braces and the empty parentheses at the end indicate that the function should be invoked with no arguments.
	}()

	// t is a pointer to a template.Template object that will be used to execute the HTML templates. The template.Template type represents a parsed template that can be executed with specific data to produce text output.
	var t *template.Template

	// If the debug variable is true, the template files are read from disk using the template.ParseFiles function. This allows developers to make changes to the template files and see their changes reflected in the output without having to recompile the binary.
	if config.Debug {
		// The template.Must function is used to create a new template.Template object by parsing the specified template files. If an error occurs while parsing the templates, the Must function panics with a descriptive error message. This line of code creates a new template.Template object by parsing the "root/layout.html" and "module1/view1.html" files and assigns it to the t variable.
		t = template.Must(template.ParseFiles("root/layout.html", "module1/view1.html"))
	} else {
		// Otherwise, if the debug variable is false, the template files are read from the embedded file system using the template.ParseFS function. This allows the binary to be distributed with all required resources embedded within it.

		// The fs.FS function is used to create a new io/fs.FS object from an embed.FS object. This allows the embedded file system to be used with functions that expect an io/fs.FS interface. This line of code creates a new io/fs.FS object from the embeddedFiles variable and assigns it to the fsys variable.
		fsys := fs.FS(embeddedFiles)

		// The template.Must function is used to create a new template.Template object by parsing the specified template files from an io/fs.FS object. If an error occurs while parsing the templates, the Must function panics with a descriptive error message. This line of code creates a new template.Template object by parsing the "root/layout.html" and "module1/view1.html" files from the fsys io/fs.FS object and assigns it to the t variable.
		t = template.Must(template.ParseFS(fsys, "root/layout.html", "module1/view1.html"))
	}

	// This statement executes the parsed templates with nil data and writes the resulting HTML page to the response.
	//
	// The t.Execute method is called on the *template.Template object that was created earlier by parsing the template files. This method takes two arguments: an io.Writer to write the resulting HTML page to, and a data value to use when executing the templates.
	//
	// In this case, we are passing the http.ResponseWriter object (named "w") as the first argument. This allows us to write the resulting HTML page directly to the response. We are also passing nil as the second argument because we do not have any data to pass to the templates at this time.
	//
	// When t.Execute is called, it will execute each of the parsed templates in turn with nil data. The resulting HTML page will be generated and written to the http.ResponseWriter object. This causes it to be sent back to the client as an HTTP response.
	t.Execute(w, nil)
}

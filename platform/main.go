package main

//Import the plugin packages
import (
	"fmt"
	_ "hajsf/gpt"
	_ "hajsf/hello"
	_ "hajsf/module1"
	_ "hajsf/module2"
	_ "hajsf/root"
	"hajsf/router"
	_ "hajsf/router"
	_ "hajsf/world"
	"net/http"
	"os"
	"os/signal"
)

//Import the pagakes required by the main function

func main() {

	router.R.Route("GET", "/static/*filepath", func(w http.ResponseWriter, r *http.Request) {

		// These wildcard parameters (asterisk *) capture all remaining parts of the URL path and store them in a map that can be
		// accessed using the URLParam function.
		// Get the file path from the URL parameter, this will not serve the subfolders
		filepath := router.URLParam(r, "filepath")

		fmt.Println("hello from route")

		// Set the Content-Type header based on the file extension
		//contentType := mime.TypeByExtension(filepath.Ext(filepath)) //  "path/filepath"
		//w.Header().Set("Content-Type", contentType)

		// Get the file path from the URL
		//filepath := strings.TrimPrefix(r.URL.Path, "/static/")

		fmt.Println("filepath:", filepath)

		// Serve the file using http.ServeFile
		http.ServeFile(w, r, "./root/static/"+filepath)
	})

	// Start the server
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", router.R)
	// Define a chanel to avoid server blocking
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

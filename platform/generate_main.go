//go:build script
// +build script

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := os.Args[1]
	pkgs := findPackagesWithRoutes(dir)
	generateMainFile(pkgs)
}

func findPackagesWithRoutes(dir string) []string {
	var pkgs []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			pkgDir := filepath.Join(dir, file.Name())
			pkgFiles, err := ioutil.ReadDir(pkgDir)
			if err != nil {
				panic(err)
			}
			for _, pkgFile := range pkgFiles {
				if strings.HasSuffix(pkgFile.Name(), ".go") {
					content, err := ioutil.ReadFile(filepath.Join(pkgDir, pkgFile.Name()))
					if err != nil {
						panic(err)
					}
					if strings.Contains(string(content), "router.R.Route") {
						pkgs = append(pkgs, filepath.ToSlash(pkgDir))
						break
					}
				}
			}
		}
	}
	return pkgs
}

func generateMainFile(pkgs []string) {
	app := "hajsf"
	f, err := os.Create("main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "//Import the plugin packages")
	fmt.Fprintln(f, "import (")
	for _, pkg := range pkgs {
		fmt.Fprintf(f, "\t_ \"%v/%v\"\n", app, pkg) // fmt.Fprintf(f, "\t_ %q\n", pkg)
	}
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "//Import the pagakes required by the main function")
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, "\t\"hajsf/router\"")
	fmt.Fprintln(f, "\t\"fmt\"")
	fmt.Fprintln(f, "\t\"net/http\"")
	fmt.Fprintln(f, "\t\"os\"")
	fmt.Fprintln(f, "\t\"os/signal\"")
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "func main() {")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "\t// Start the server")
	fmt.Fprintln(f, "\tfmt.Println(\"Starting server...\")")
	fmt.Fprintln(f, "\thttp.ListenAndServe(\":8080\", router.R)")
	fmt.Fprintln(f, "\t// Define a chanel to avoid server blocking")
	fmt.Fprintln(f, "\tc := make(chan os.Signal, 1)")
	fmt.Fprintln(f, "\tsignal.Notify(c, os.Interrupt)")
	fmt.Fprintln(f, "\t<-c")
	fmt.Fprintln(f, "}")
}

/*
//dynamically load packages that define routes without having to modify main.go is to use code generation.

// For example, you could write a script that scans a specific directory for packages
// that define routes and generates a Go file that imports all those packages.

// This script takes as an argument the path to the directory containing
// the packages that define routes and generates a main.go file that
// imports all those packages. You could run this script as part of your
// build process to automatically generate the main.go file.


To compile and build the script I provided in my previous message, you can save it to a file with a .go extension (e.g. generate_main.go) and then run the go build command to build it:

go build -tags script generate_main.go

This will create an executable file named generate_main (or generate_main.exe on Windows) that you can run to generate the main.go file:

./generate_main path/to/packages

Make sure to replace path/to/packages with the actual path to the directory containing the packages that define routes.

Once you have generated the main.go file, you can build your application as usual by running the go build command in the directory containing your application’s source code.
*/

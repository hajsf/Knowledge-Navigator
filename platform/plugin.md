If you want to avoid having to manually import each package that defines routes in main.go, one approach could be to use Go’s plugin package to dynamically load the packages at runtime.

For example, you could compile each package that defines routes into a plugin by adding a file like this to each package:
```go
package main

import _ "hajsf/hello"

func main() {}
```

And then build the plugin with a command like this:
```bash
go build -buildmode=plugin -o hello.so
```
Then in your main.go file you could use the plugin package to dynamically load all the plugins from a specific directory:
```go
package main

import (
	"io/ioutil"
	"path/filepath"
	"plugin"

	"hajsf/router"
)

func main() {
	r := router.New()
	loadPlugins(r)
	// ...
}

func loadPlugins(r *router.Router) {
	files, err := ioutil.ReadDir("plugins")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".so" {
			p, err := plugin.Open(filepath.Join("plugins", file.Name()))
			if err != nil {
				panic(err)
			}
			addRoutesSym, err := p.Lookup("AddRoutes")
			if err != nil {
				panic(err)
			}
			addRoutesFunc, ok := addRoutesSym.(func(*router.Router))
			if !ok {
				panic("AddRoutes symbol is not of type func(*router.Router)")
			}
			addRoutesFunc(r)
		}
	}
}
```
This way you can add new packages that define routes without having to modify main.go. However, keep in mind that using plugins can make your code more complex and harder to test.
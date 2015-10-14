// main.go
package main

import (
	"html/template"
	"strings"

	"github.com/go-macaron/jade"
	"gopkg.in/macaron.v1"
)

func FuncUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	m := macaron.Classic()
	// render html templates from templates directory
	m.Use(jade.Renderer(jade.Options{
		Directory:  "templates",       // Specify what path to load the templates from.
		Extensions: []string{".jade"}, // Specify extensions to load for templates.
		Funcs: []template.FuncMap{map[string]interface{}{
			"upper": FuncUpper,
		}}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))

	m.Get("/", func(r jade.Render) {
		r.HTML(200, "hello", map[string]string{
			"foo": "bar",
		})
	})

	m.Run()
}

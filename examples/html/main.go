// main.go
package main

import (
	"github.com/go-macaron/jade"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	// render html templates from templates directory
	m.Use(jade.Renderer())

	m.Get("/", func(r jade.Render) {
		r.HTML(200, "hello", map[string]string{
			"foo": "bar",
		})
	})

	m.Run()
}

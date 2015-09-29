// main.go
package main

import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/jade"
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

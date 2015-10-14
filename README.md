# Macaron-jade
[![GoDoc](https://godoc.org/github.com/go-macaron/jade?status.svg)](https://godoc.org/github.com/go-macaron/jade)

Macaron middleware/handler for easily rendering serialized JSON and HTML template responses from Jade templates.

If you donot know about jade, [learn from here](http://www.learnjade.com/)

## Usage
This middleware uses Jade implementation in Go [go-floki/jade](https://github.com/go-floki/jade) to render Jade templates.

Some examples can be found in [examples](examples)

~~~ go
// main.go
package main

import (
  "gopkg.in/macaron.v1"
	"github.com/go-macaron/jade"
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
~~~

File `templates/hello.jade`

~~~ html
h2 Hello #{foo}!
~~~

### Options
`jade.Renderer` comes with a variety of configuration options:

**Layout** is not supported.

~~~ go
// ...
m.Use(jade.Renderer(jade.Options{
  Directory: "templates", // Specify what path to load the templates from.
  Extensions: []string{".jade"}, // Specify extensions to load for templates.
  Funcs: []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
  Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
  IndentJSON: true, // Output human readable JSON
}))
// ...
~~~

### Loading Templates
By default the `jade.Renderer` middleware will attempt to load templates with a '.jade' extension from the "templates" directory. Templates are found by traversing the templates directory and are named by path and basename. For instance, the following directory structure:

~~~
templates/
  |
  |__ admin/
  |      |
  |      |__ index.jade
  |      |
  |      |__ edit.jade
  |
  |__ home.jade
~~~

Will provide the following templates:
~~~
admin/index
admin/edit
home
~~~

### Character Encodings
The `jade.Renderer` middleware will automatically set the proper Content-Type header based on which function you call. See below for an example of what the default settings would output (note that UTF-8 is the default):
~~~ go
// main.go
package main

import (
  "gopkg.in/macaron.v1"
  "github.com/go-macaron/jade"
)

func main() {
  m := macaron.Classic()
  m.Use(jade.Renderer())

  // This will set the Content-Type header to "text/html; charset=UTF-8"
  m.Get("/", func(r jade.Render) {
    r.HTML(200, "hello", "world")
  })

  // This will set the Content-Type header to "application/json; charset=UTF-8"
  m.Get("/api", func(r jade.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Run()
}

~~~

In order to change the charset, you can set the `Charset` within the `jade.Options` to your encoding value:
~~~ go
// main.go
package main

import (
  "gopkg.in/macaron.v1"
  "github.com/go-macaron/jade"
)

func main() {
  m := macaron.Classic()
  m.Use(jade.Renderer(render.Options{
    Charset: "ISO-8859-1",
  }))

  // This is set the Content-Type to "text/html; charset=ISO-8859-1"
  m.Get("/", func(r jade.Render) {
    r.HTML(200, "hello", "world")
  })

  // This is set the Content-Type to "application/json; charset=ISO-8859-1"
  m.Get("/api", func(r jade.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Run()
}

~~~

## Authors
* [Jeremy Saenz](http://github.com/codegangsta)
* [Cory Jacobsen](http://github.com/cojac)
* [frogprog](http://github.com/frogprog)
* [codeskyblue](http://github.com/codeskyblue)
* [Unknwon](http://github.com/Unknwon)

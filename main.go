package main

import (
	"github.com/go-martini/martini"
)

//https://github.com/go-martini/martini
//https://github.com/go-martini/martini
func main() {
	m := martini.Classic()
	m.Use(martini.Static("public"))
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/hello/:name", func(params martini.Params) string {
		return "hello " + params["name"]
	})
	m.NotFound(func() (int, string) {
		return 404, "NOT FOUND !!"
	})
	m.RunOnAddr(":8000")
}

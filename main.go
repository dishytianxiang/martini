package main

import (
	"github.com/go-martini/martini"
)

//https://github.com/go-martini/martini
func main() {
	m := martini.Classic()
	m.Get("/",functioin(){
		return "Hello world!"
	})
}

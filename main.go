package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

//https://github.com/go-martini/martini
//https://github.com/go-martini/martini
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("public"))
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/login", func(r render.Render) {

		r.HTML(200, "login", nil)
	})
	m.Get("/login/:userName", func(params martini.Params) string {
		if params["userName"] == "dishy" {
			fmt.Println(params["userName"])
		}
		return "hello " + params["userName"]
	})
	m.NotFound(func() (int, string) {
		return 404, "NOT FOUND !!"
	})
	m.RunOnAddr(":8000")
}
func setCookie() {
	//expiration := time.Now()
	//expiration = expiration.AddDate(1, 0, 0)
	//cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	//http.SetCookie(w, &cookie)
}

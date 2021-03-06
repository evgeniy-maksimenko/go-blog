package routes

import (
	"../session"
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func LogoutHandler(rnd render.Render, r *http.Request, s *session.Session) {
	s.Username = ""
	s.IsAuth = false

	rnd.Redirect("/")
}

func PostLoginHandler(rnd render.Render, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(password)

	s.Username = username
	s.IsAuth = true

	rnd.Redirect("/")
}

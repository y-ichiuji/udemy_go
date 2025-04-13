package controllers

import (
	"log"
	"net/http"

	"github.com/y-ichiuji/udemy-go/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.Create(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}
package handlers

import (
	"fmt"
	"net/http"
	"pagetest/go/Users"
	"text/template"
)

var user Users.User

func CssHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got css request")
	w.Header().Set("Content-type", "text/css")

	filename := r.URL.Path[len("/fonts/style/"):]

	http.ServeFile(w, r, "./fonts/style/"+filename)
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got js request")
	w.Header().Set("Content-Type", "application/javascript")

	filename := r.URL.Path[len("/fonts/scripts/"):]
	http.ServeFile(w, r, "./fonts/scripts/"+filename)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")

	t, err := template.ParseFiles("./static/hello.html")

	if err != nil {
		fmt.Println("error on the parsing phase")
	}

	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/login.html")
		t.Execute(w, nil)
	}

	if r.Method == "POST" {
		user.Name = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.ValidatePassword()
		if user.ValidPass {
			fmt.Println("Your username is ", user.Name)
			fmt.Println("Your password is", user.Password)
		}

		t, _ := template.ParseFiles("static/login.html")
		t.Execute(w, nil)
	}
}

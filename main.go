package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const SERVERPORT string = ":8000"

var user User

type User struct {
	Name      string
	Password  string
	ValidPass bool
}

func (u User) validatePassword() {
	if len(u.Password) > 8 {
		u.ValidPass = true
		fmt.Println("User ", user.Name)
		fmt.Println("Password ", user.Password)

	} else {
		u.ValidPass = false
		fmt.Println("invalid Password, it must contain more than 8 characters")

	}

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	t, _ := template.ParseFiles("./static/hello.html")
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/login.html")
		t.Execute(w, nil)
	} else {
		user.Name = r.FormValue("username")
		user.Password = r.FormValue("password")
        user.validatePassword()

		t, _ := template.ParseFiles("static/login.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running at port", SERVERPORT)

	http.ListenAndServe(SERVERPORT, nil)
}

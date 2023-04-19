package main

import (
	"fmt"
	"net/http"
	User "pagetest/m/v2/User"
	"text/template"
)

const SERVERPORT string = ":8000"

var user User.User

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
		user.ValidatePassword()
		if user.ValidPass {
			fmt.Println("Your username is ", user.Name)
			fmt.Println("Your password is", user.Password)
		}

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

package main

import (
	"fmt"
	"net/http"
	User "pagetest/go/User"
	"text/template"
)

const SERVERPORT string = ":8000"

var user User.User

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println(fs)
	// mime.AddExtensionType(".js", "application/javascript")
	// mime.AddExtensionType(".css", "text/css")

	// File handlers for javascript, css.
	http.HandleFunc("/fonts/style/", cssHandler)
	http.HandleFunc("/fonts/scripts/", jsHandler)

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running at port", SERVERPORT)
	http.ListenAndServe(SERVERPORT, nil)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got css request")
	w.Header().Set("Content-type", "text/css")

	filename := r.URL.Path[len("/fonts/style/"):]

	http.ServeFile(w, r, "./fonts/style/"+filename)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got js request")
	w.Header().Set("Content-Type", "application/javascript")

	filename := r.URL.Path[len("/fonts/scripts/"):]
	fmt.Println(filename)

	http.ServeFile(w, r, "./fonts/scripts/"+filename)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	t, err := template.ParseFiles("./static/hello.html")

	if err != nil {
		fmt.Println("error on the parsing phase")
	}

	t.Execute(w, nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/login.html")
		t.Execute(w, nil)
	} else {
		user.Name = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.ValidatePassword()
		if user.ValidPass {
			fmt.Println("Your username is ", user.Name)
			fmt.Println("Your password is", user.Password)
		}

		t, _ := template.ParseFiles("./static/login.html")
		t.Execute(w, nil)
	}
}

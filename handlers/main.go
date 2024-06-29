package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	Users "pagetest/go/users"
	"path"
	"text/template"
)

var user Users.User

type StatusJson struct {
	GolangVersion string `json:"golangVersion"`
	Working       bool   `json:"working"`
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got css request")
	w.Header().Set("Content-type", "text/css")

	filename := r.URL.Path[len("/fonts/style/"):]
	filename = path.Clean(filename)

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
		fmt.Println("Error while parsing: ", err)
	}

	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./static/login.html")
		if err != nil {
			fmt.Println("Error while parsing: ", err)
		}
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

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	statusJson := &StatusJson{
		GolangVersion: "1.22",
	}

	if statusJson.GolangVersion == "1.22" {
		statusJson.Working = true
	}

	err := json.NewDecoder(r.Body).Decode(statusJson)
	if err != nil {
		fmt.Printf("Error decoding json, Error: %v", err)
	}

	statusDecoded, err := json.Marshal(statusJson)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(statusDecoded)

	// w.Header().Set("Content-type", "application/json")
	// err := json.NewEncoder(w).Encode(&StatusJson{
	// 	GolangVersion: "1.22",
	// 	Working:       "yes",
	// })

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	status = 404
	// }
	//
	// w.WriteHeader(status)

}

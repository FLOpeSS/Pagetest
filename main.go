package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const SERVERPORT string = ":8000"

type Page struct {
    Title string
    link string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got / request")
    http.ServeFile(w, r, "./static/index.html")
}


func secondHandler(w http.ResponseWriter, r *http.Request) {
    for i := 0; i < 100; i++ {
        fmt.Println(i)
    }
    tmpl, _ := template.ParseFiles("static/index.html")
    data := Page{
        Title: "hello",
        link : "main.js",
    }
    tmpl.Execute(w, data)
}


func main() {
    http.HandleFunc("/", homePageHandler)
    fmt.Println("Server running at port", SERVERPORT)
    http.ListenAndServe(SERVERPORT, nil)
}






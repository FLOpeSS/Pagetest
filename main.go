package main

import (
    "fmt"
    "net/http"
)

const SERVERPORT = ":8000"

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("got / request")
    http.ServeFile(w, r, "index.html")
}


func main() {
    http.HandleFunc("/", homePageHandler)
    http.ListenAndServe(SERVERPORT, nil)
}


package main

import (
	"fmt"
	"net/http"
	"pagetest/go/handlers"
)

const SERVERPORT string = ":3000"

func main() {

	// File handlers for javascript, css.
	http.HandleFunc("/fonts/scripts/", handlers.JsHandler)
	http.HandleFunc("/fonts/style/", handlers.CssHandler)

	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	http.HandleFunc("/api/v1/status", handlers.StatusHandler)

	fmt.Println("Server running at port", SERVERPORT)

	err := http.ListenAndServe(SERVERPORT, nil)
	if err != nil {
		fmt.Println("Failed to start the server", err)
	}
}

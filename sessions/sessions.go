package sessions

import (
    "fmt"
    "net/http"
)


func fetchCookie(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("username")
    fmt.Prinln(w, cookie)
}



package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }

        if r.Method == "GET" {
            fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
        } else if r.Method == "POST" {
            fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
        } else {
            http.Error(w, "Invalid request method.", 405)
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

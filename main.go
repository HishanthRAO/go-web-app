package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello from DevOpsified Go Web App!</h1><p>Built by Hishanth Rao</p>")
    })
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}

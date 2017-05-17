package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<html><head /><body style='background-color: yellow'>Hello Travis</body></html>!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":80", nil)
}


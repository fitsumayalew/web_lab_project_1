package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

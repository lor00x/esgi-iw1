package main

import (
	"embed"
	"net/http"
)

//go:embed public
var public embed.FS

func main() {
	// http.ServeMux
	s := http.NewServeMux()

	// http.FileSystem
	f := http.FS(public)

	// http.Handler
	h := http.FileServer(f)

	s.Handle("/public/", h)

	http.ListenAndServe(":8080", s)
}

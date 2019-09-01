package main

import (
	"net/http"

	_ "./statik"

	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, _ := fs.New()
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}

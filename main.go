package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/kojisaiki/webappsample01/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	fmt.Println("start")

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("new")

	f, err := statikFS.Open("/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("opened")
	defer f.Close()
	io.Copy(os.Stdout, f)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

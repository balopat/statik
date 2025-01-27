//go:generate go run ./statik/gen/gen.go -src=./public

package main

import (
	"log"
	"net/http"

	_ "github.com/balopat/statik/example/statik"
	"github.com/balopat/statik/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}

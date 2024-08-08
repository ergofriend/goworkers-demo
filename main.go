package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/syumai/workers"
)

func main() {
	http.Handle("/", templ.Handler(hello("World")))
	workers.Serve(nil) // use http.DefaultServeMux
}

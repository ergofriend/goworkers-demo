package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/syumai/workers"
)

//go:embed templates/*
var templates embed.FS

func main() {
	tmpl, err := template.ParseFS(templates, "templates/*")
	if err != nil {
		log.Fatalf("template parse error: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err = tmpl.ExecuteTemplate(w, "hello.go.tmpl", map[string]interface{}{"Name": "World"})
		if err != nil {
			log.Fatalf("template execute error: %v", err)
			return
		}
	})

	workers.Serve(nil) // use http.DefaultServeMux
}

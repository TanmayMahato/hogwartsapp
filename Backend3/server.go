package main

import (
	"html/template"

	"log"
	"net/http"
)

type Server struct {
	ServerName string
}

func main() {

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("Frontend/index.html"))
		//name := Server{ServerName: "server3"}

		tmpl.Execute(w, nil)

	}

	http.HandleFunc("/", handler1)

	log.Fatal(http.ListenAndServe(":8003", nil))
}

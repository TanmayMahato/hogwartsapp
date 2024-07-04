package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	f := "http://localhost:5004"
	targetURL, _ := url.Parse(f)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.Handle("/", proxy)

	log.Fatal(http.ListenAndServe(":8004", nil))
}

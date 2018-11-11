package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var ver string
	ver = "1.6"
	name, _ := os.Hostname()
	fmt.Fprintf(w, "<html><head><title>Test site</title></head>")
	fmt.Fprintf(w, "<h1>This is a simple golang test app</h1>")
	fmt.Fprintf(w, "<p>Server by hostname: %s</p>", name)
	fmt.Fprintf(w, "<p>Version: %s</p>", ver)
	fmt.Fprintf(w, "<img src=\"img/gradient.png\">")
	fmt.Fprintf(w, "</body></html>")
}

func main() {
	images := http.FileServer(http.Dir("./img"))
	http.HandleFunc("/", handler)
	http.Handle("/img/", http.StripPrefix("/img/", images))
	http.ListenAndServe(":8080", nil)
}

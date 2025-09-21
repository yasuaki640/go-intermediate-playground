package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello")
	}
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "post article dummy")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
